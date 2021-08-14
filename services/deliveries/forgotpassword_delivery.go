package deliveries

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/mailgun/mailgun-go/v4"
)

type (
	forgotPasswordHandler struct {
		forgotPasswordUseCase domains.ForgorPasswordUseCase
	}
	paramEmail struct {
		Email string `json:"email" validate:"required,email,min=6,max=255"`
	}
	templateResetPassword struct {
		AppName  string
		LinkHash string
	}
)

func NewForgotPasswordHandler(usecase domains.ForgorPasswordUseCase, u *[]models.ModuleMethodSlug) *forgotPasswordHandler {
	moduleRepo := repositories.NewModuleRepository(databases.DB)
	moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
	for _, value := range *u {
		newModule := models.Modules{
			Name:        value.Name,
			Description: value.Description,
			ModuleSlug:  value.Slug,
			Method:      value.Method,
		}
		err := moduleUseCase.GetModuleBySlug(&newModule, value.Method, value.Slug)
		if err != nil {
			moduleUseCase.CreateModule(&newModule)
		}
	}
	return &forgotPasswordHandler{
		forgotPasswordUseCase: usecase,
	}
}
func (f *forgotPasswordHandler) ResetPassword(c *fiber.Ctx) error {
	var param paramEmail
	err := c.BodyParser(&param)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	linkHash, err := f.forgotPasswordUseCase.CreateForgotPasswordHash(param.Email)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_HASHLINK_IN_REDIS, fiber.StatusBadRequest)
	}
	mg := mailgun.NewMailgun(config.ENV.MyDomain, config.ENV.MailGunApiKey)
	sender := config.ENV.EmailAdministrator
	subject := config.ENV.AppName + " reset password"
	recipient := param.Email
	message := mg.NewMessage(sender, subject, "", recipient)
	t, _ := template.ParseFiles("../views/change_password/change_password.html")
	tplData := templateResetPassword{
		AppName:  config.ENV.AppName,
		LinkHash: linkHash,
	}
	var tpl bytes.Buffer
	if err := t.Execute(&tpl, tplData); err != nil {
		return helpers.FailOnError(c, err, constants.ERR_TEMPLATE_PARSE_DATA, fiber.StatusNotAcceptable)
	}
	body := tpl.String()

	message.SetHtml(body)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	resp, id, err := mg.Send(ctx, message)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_MAIL_GUN_FORBIDDEN, fiber.StatusUnauthorized)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_SEND_LINK_RESET_PASSWORD_TO_EMAIL_SUCCESSFUL,
		"error": nil,
		"id":    id,
		"data":  resp,
	})
}

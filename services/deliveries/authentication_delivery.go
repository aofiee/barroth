package deliveries

import (
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
)

type (
	authenticationHandler struct {
		authenticationUseCase domains.AuthenticationUseCase
		moduleName            string
		description           string
		slug                  string
	}
	paramsLogin struct {
		Email    string `json:"email" validate:"required,email,min=6,max=255"`
		Password string `json:"password" validate:"required,min=6,max=64"`
	}
)

func NewAuthenHandler(usecase domains.AuthenticationUseCase, m, d, u string) *authenticationHandler {
	newModule := models.Modules{
		Name:        m,
		Description: d,
		ModuleSlug:  u,
	}
	moduleRepo := repositories.NewModuleRepository(databases.DB)
	moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
	err := moduleUseCase.GetModule(&newModule, u)
	if err != nil {
		moduleUseCase.CreateModule(&newModule)
	}
	return &authenticationHandler{
		authenticationUseCase: usecase,
		moduleName:            m,
		description:           d,
		slug:                  u,
	}
}

func (a *authenticationHandler) Login(c *fiber.Ctx) error {
	var auth paramsLogin
	err := c.BodyParser(&auth)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&auth)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	var u models.Users
	err = a.authenticationUseCase.Login(&u, auth.Email, auth.Password)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_USERNAME_PASSWORD_INCORRECT, fiber.StatusUnauthorized)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_LOGIN_SUCCESSFUL,
		"error": nil,
	})
}

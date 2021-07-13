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
	userHandler struct {
		userUseCase domains.UserUseCase
		moduleName  string
		description string
		slug        string
	}
	paramsUser struct {
		Email     string `json:"email" form:"email" validate:"required,email,min=6,max=255"`
		Password  string `json:"password" form:"password" validate:"required,min=6,max=64"`
		Telephone string `json:"telephone" form:"telephone" validate:"required,min=10,max=50"`
		Name      string `json:"name" form:"name" validate:"required,min=6,max=255"`
		RoleID    int    `json:"role_id" form:"role_id" validate:"required,number"`
	}
)

func NewUserHandelr(usecase domains.UserUseCase, m, d, u string) *userHandler {
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
	return &userHandler{
		userUseCase: usecase,
		moduleName:  m,
		description: d,
		slug:        u,
	}
}
func (u *userHandler) NewUser(c *fiber.Ctx) error {
	var nu paramsUser
	err := c.BodyParser(&nu)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&nu)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	user := models.Users{
		Email:     nu.Email,
		Password:  nu.Password,
		Name:      nu.Name,
		Telephone: nu.Telephone,
	}
	err = u.userUseCase.CreateUser(&user)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_CREATE_USER_SUCCESSFUL,
		"error": nil,
	})
}

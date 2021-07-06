package deliveries

import (
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
)

type (
	roleHandler struct {
		roleUseCase domains.RoleUseCase
		moduleName  string
		description string
		slug        string
	}
)

func NewRoleHandelr(usecase domains.RoleUseCase, m, d, u string) *roleHandler {
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
	return &roleHandler{
		roleUseCase: usecase,
		moduleName:  m,
		description: d,
		slug:        u,
	}
}
func (r *roleHandler) NewRole(c *fiber.Ctx) error {
	var role models.RoleItems
	err := c.BodyParser(&role)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot parse json", fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&role)
	if errorResponse != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   "input error.",
			"error": errorResponse,
		})
	}
	err = r.roleUseCase.CreateRole(&role)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot create role", fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "create role successful.",
		"error": nil,
	})
}

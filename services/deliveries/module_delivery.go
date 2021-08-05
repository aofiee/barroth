package deliveries

import (
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
)

type (
	moduleHandler struct {
		moduleUseCase domains.ModuleUseCase
	}
)

func NewModuleHandler(usecase domains.ModuleUseCase, u *[]models.ModuleMethodSlug) *moduleHandler {
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
	return &moduleHandler{
		moduleUseCase: usecase,
	}
}
func (m *moduleHandler) GetAllModules(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_GET_ALL_ROLE_SUCCESSFULE,
		"error": nil,
		"data":  nil,
	})
}

package deliveries

import (
	"time"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

type (
	systemHandler struct {
		systemUseCase domains.SystemUseCase
		moduleName    string
		description   string
	}
)

func NewSystemHandelr(usecase domains.SystemUseCase, m, d string, u *[]models.ModuleMethodSlug) *systemHandler {
	for _, value := range *u {
		newModule := models.Modules{
			Name:        m,
			Description: d,
			ModuleSlug:  value.Slug,
			Method:      value.Method,
		}
		moduleRepo := repositories.NewModuleRepository(databases.DB)
		moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
		err := moduleUseCase.GetModuleBySlug(&newModule, value.Method, value.Slug)
		if err != nil {
			moduleUseCase.CreateModule(&newModule)
		}
	}
	return &systemHandler{
		systemUseCase: usecase,
		moduleName:    m,
		description:   d,
	}
}
func (s *systemHandler) SystemInstallation(c *fiber.Ctx) error {
	var systems models.System
	err := s.systemUseCase.GetFirstSystemInstallation(&systems)
	if err != nil {
		systems.AppName = barroth_config.ENV.AppName
		systems.SiteURL = barroth_config.ENV.SiteURL + ":" + barroth_config.ENV.AppPort
		systems.IsInstall = 0
		systems.CreatedAt = time.Now()
		err = s.systemUseCase.CreateSystem(&systems)
		if err != nil {
			return helpers.FailOnError(c, err, "can not create new record", fiber.StatusBadRequest)
		}

		user := models.Users{
			Email:     barroth_config.ENV.EmailAdministrator,
			Password:  barroth_config.ENV.PasswordAdministrator,
			Telephone: barroth_config.ENV.TelephoneAdministrator,
			UUID:      utils.UUIDv4(),
		}
		role := models.RoleItems{
			Name:        "Administrator",
			Description: "Initial Role",
		}
		err = s.systemUseCase.CreateRole(&role)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
		}
		user.UserRoleID.RoleItemID = role.ID
		err := s.systemUseCase.CreateUser(&user)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg":   "complete the installation.",
			"error": nil,
		})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"msg":   "software is already installed.",
		"error": nil,
	})
}

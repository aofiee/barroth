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
	}
)

func NewSystemHandelr(usecase domains.SystemUseCase, u *[]models.ModuleMethodSlug) *systemHandler {
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
	return &systemHandler{
		systemUseCase: usecase,
	}
}
func (s *systemHandler) SystemInstallation(c *fiber.Ctx) error {
	var systems models.System
	err := s.systemUseCase.GetFirstSystemInstallation(&systems)
	if err != nil {
		user := models.Users{
			Email:     barroth_config.ENV.EmailAdministrator,
			Password:  barroth_config.ENV.PasswordAdministrator,
			Telephone: barroth_config.ENV.TelephoneAdministrator,
			UUID:      utils.UUIDv4(),
			Provider:  "system",
		}
		role := models.RoleItems{
			Name:        "Administrator",
			Description: "Initial Role",
		}
		err = s.systemUseCase.CreateRole(&role)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
		}
		var modules []models.Modules
		err := s.systemUseCase.SetExecToAllModules(&modules, role.ID, 1)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_SET_EXEC_ALL_MODULE, fiber.StatusBadRequest)
		}
		user.UserRoleID.RoleItemID = role.ID
		err = s.systemUseCase.CreateUser(&user)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_USER, fiber.StatusBadRequest)
		}
		systems.AppName = barroth_config.ENV.AppName
		systems.SiteURL = barroth_config.ENV.SiteURL + ":" + barroth_config.ENV.AppPort
		systems.IsInstall = 1
		systems.CreatedAt = time.Now()
		err = s.systemUseCase.CreateSystem(&systems)
		if err != nil {
			return helpers.FailOnError(c, err, "can not create new record", fiber.StatusBadRequest)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg":   "complete the installation",
			"error": nil,
		})
	}
	return c.Status(fiber.StatusFound).JSON(fiber.Map{
		"msg":   "software is already installed.",
		"error": nil,
	})
}

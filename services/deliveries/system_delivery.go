package deliveries

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	fiber "github.com/gofiber/fiber/v2"
)

type (
	systemHandler struct {
		systemUseCase domains.SystemUseCase
		moduleName    string
		description   string
	}
)

func NewSystemHandelr(usecase domains.SystemUseCase, m, d string) *systemHandler {
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
		err = s.systemUseCase.CreateSystem(&systems)
		if err != nil {
			return helpers.FailOnError(c, err, "can not create new record", fiber.StatusBadRequest)
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

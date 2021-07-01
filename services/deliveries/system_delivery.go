package deliveries

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	fiber "github.com/gofiber/fiber/v2"
)

type (
	systemHandler struct {
		systemUseCase domains.SystemUseCase
	}
)

func NewSystemHandelr(usecase domains.SystemUseCase) *systemHandler {
	return &systemHandler{
		systemUseCase: usecase,
	}
}
func (s *systemHandler) SystemInstallation(c *fiber.Ctx) error {
	var systems models.System
	err := s.systemUseCase.GetFirstSystemInstallation(&systems)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot parse json", fiber.StatusBadRequest)
	}
	// err = c.BodyParser(&systems)
	// if err != nil {
	// 	return helpers.FailOnError(c, err, "cannot parse json", fiber.StatusBadRequest)
	// }
	return nil
}

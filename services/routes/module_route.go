package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/models"
	"github.com/gofiber/fiber/v2"
)

type (
	moduleRoutes struct {
		config barroth_config.Config
	}
)

func NewModuleRoutes(config barroth_config.Config) *moduleRoutes {
	return &moduleRoutes{
		config: config,
	}
}
func (m *moduleRoutes) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Name:        "Get All Modules",
			Description: "ดึงรายการ Module ทั้งหมดที่มีในระบบ",
			Method:      fiber.MethodGet,
			Slug:        RoleSlug,
		},
	)
	// repo := repositories.NewModuleRepository(databases.DB)
	// u := usecases.NewModuleUseCase(repo)
	// handler := deliveries.NewModuleHandler(u, &moduleRoute)

}

package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/deliveries"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
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
			Slug:        "/modules",
		},
		models.ModuleMethodSlug{
			Name:        "Update Modules",
			Description: "แก้ไขรายละเอียด Modules",
			Method:      fiber.MethodPut,
			Slug:        "/module",
		},
	)
	repo := repositories.NewModuleRepository(databases.DB)
	u := usecases.NewModuleUseCase(repo)
	handler := deliveries.NewModuleHandler(u, &moduleRoute)
	e := app.Group("/module", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Put("/:id", handler.UpdateModule)

	e = app.Group("/modules", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Get("/", handler.GetAllModules)
}

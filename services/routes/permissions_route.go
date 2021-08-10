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
	permissionsRoutes struct {
		config barroth_config.Config
	}
)

func NewPermissionRoutes(config barroth_config.Config) *permissionsRoutes {
	return &permissionsRoutes{
		config: config,
	}
}
func (m *permissionsRoutes) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Name:        "Set Permissions to Role",
			Description: "ตั้งค่า permission ให้กับ role ต่างๆ",
			Method:      fiber.MethodPut,
			Slug:        "/permissions",
		},
	)
	repo := repositories.NewPermissionsRepository(databases.DB)
	u := usecases.NewPermissionsUseCase(repo)
	handler := deliveries.NewPermissionHandler(u, &moduleRoute)
	e := app.Group("/permissions", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Put("/:id", handler.SetPermission)
}

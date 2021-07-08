package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/deliveries"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
)

type (
	roleRoutes struct {
		config barroth_config.Config
	}
)

func NewRoleRoutes(config barroth_config.Config) *roleRoutes {
	return &roleRoutes{
		config: config,
	}
}
func (r *roleRoutes) Install(app *fiber.App) {
	repo := repositories.NewRoleRepository(databases.DB)
	u := usecases.NewRoleUseCase(repo)
	handler := deliveries.NewRoleHandelr(u, "Installation", "Installation Module This is an API group for the system installation environment.", "/role")
	e := app.Group("/role")
	e.Post("/", handler.NewRole)
	e.Get("/:id", handler.GetRole)
	e.Put("/:id", handler.UpdateRole)

	e = app.Group("/roles")
	e.Get("/", handler.GetAllRoles)
	e.Delete("/", handler.DeleteRoles)
	e.Put("/restore", handler.RestoreRoles)
}

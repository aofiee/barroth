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
	roleRoutes struct {
		config barroth_config.Config
	}
)

const (
	RoleSlug         = "/role"
	RolesSlug        = "/roles"
	RestoreRolesSlug = "/roles/restore"
)

func NewRoleRoutes(config barroth_config.Config) *roleRoutes {
	return &roleRoutes{
		config: config,
	}
}
func (r *roleRoutes) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Method: fiber.MethodPost,
			Slug:   RoleSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodGet,
			Slug:   RoleSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodPut,
			Slug:   RoleSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodGet,
			Slug:   RolesSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodDelete,
			Slug:   RolesSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodPut,
			Slug:   RestoreRolesSlug,
		},
	)
	repo := repositories.NewRoleRepository(databases.DB)
	u := usecases.NewRoleUseCase(repo)
	handler := deliveries.NewRoleHandelr(u, "Role", "Manage Role Module", &moduleRoute)

	authRepo := repositories.NewAuthenticationRepository(databases.DB, databases.QueueClient)
	authUseCase := usecases.NewAuthenticationUseCase(authRepo)
	authHandler := deliveries.GetAuthHandlerUsecase(authUseCase)

	e := app.Group("/role", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken)
	e.Post("/", handler.NewRole)
	e.Get("/:id", handler.GetRole)
	e.Put("/:id", handler.UpdateRole)

	e = app.Group("/roles", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken)
	e.Get("/", handler.GetAllRoles)
	e.Delete("/", handler.DeleteRoles)
	e.Put("/restore", handler.RestoreRoles)
}

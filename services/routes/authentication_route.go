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
	authenticationRoutes struct {
		config barroth_config.Config
	}
)

func NewAuthenticationRoutes(config barroth_config.Config) *authenticationRoutes {
	return &authenticationRoutes{
		config: config,
	}
}
func (r *authenticationRoutes) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Method: fiber.MethodPost,
			Slug:   "/auth",
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodDelete,
			Slug:   "/auth/logout",
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodPost,
			Slug:   "/auth/refresh_token",
		},
	)
	repo := repositories.NewAuthenticationRepository(databases.DB, databases.QueueClient)
	u := usecases.NewAuthenticationUseCase(repo)
	handler := deliveries.NewAuthenHandler(u, "Authentication", "Authentication Module", &moduleRoute)
	e := app.Group("/auth")
	e.Post("/", handler.Login)
	e.Delete("/logout", handler.AuthorizationRequired(), handler.Logout)
	e.Post("/refresh_token", handler.RefreshToken)
}

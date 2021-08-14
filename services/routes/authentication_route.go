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
			Name:        "Login",
			Description: "เข้าสู่ระบบเพื่อขอ Token",
			Method:      fiber.MethodPost,
			Slug:        "/auth",
		},
		models.ModuleMethodSlug{
			Name:        "Logout",
			Description: "ออกจากระบบ",
			Method:      fiber.MethodDelete,
			Slug:        "/auth/logout",
		},
		models.ModuleMethodSlug{
			Name:        "Refresh Token",
			Description: "ขอ Access Token ใหม่ ด้วย Refresh Token",
			Method:      fiber.MethodPost,
			Slug:        "/auth/refresh_token",
		},
	)
	repo := repositories.NewAuthenticationRepository(databases.DB, databases.TokenQueueClient)
	u := usecases.NewAuthenticationUseCase(repo)
	handler := deliveries.NewAuthenHandler(u, &moduleRoute)
	e := app.Group("/auth")
	e.Post("/", handler.Login)
	e.Delete("/logout", handler.AuthorizationRequired(), handler.IsRevokeToken, handler.Logout)
	e.Post("/refresh_token", handler.RefreshToken)
}

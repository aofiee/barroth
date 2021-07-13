package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/deliveries"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	fiber "github.com/gofiber/fiber/v2"
)

type (
	userRoutes struct {
		config barroth_config.Config
	}
)

func NewUserRoutes(config barroth_config.Config) *userRoutes {
	return &userRoutes{
		config: config,
	}
}
func (r *userRoutes) Install(app *fiber.App) {
	repo := repositories.NewUserRepository(databases.DB)
	u := usecases.NewUserUseCase(repo)
	handler := deliveries.NewUserHandelr(u, "Users", "Installation Module This is an API group for the system installation environment.", "/user")
	e := app.Group("/user")
	e.Post("/", handler.NewUser)
	// e = app.Group("/users")
}

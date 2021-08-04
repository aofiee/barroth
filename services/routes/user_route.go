package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/deliveries"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	fiber "github.com/gofiber/fiber/v2"
)

type (
	userRoutes struct {
		config barroth_config.Config
	}
)

const (
	UserSlug  = "/user"
	UsersSlug = "/users"
)

func NewUserRoutes(config barroth_config.Config) *userRoutes {
	return &userRoutes{
		config: config,
	}
}
func (r *userRoutes) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Method: fiber.MethodPost,
			Slug:   UserSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodPut,
			Slug:   UserSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodGet,
			Slug:   UserSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodDelete,
			Slug:   UserSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodDelete,
			Slug:   UsersSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodGet,
			Slug:   UsersSlug,
		},
		models.ModuleMethodSlug{
			Method: fiber.MethodPut,
			Slug:   UsersSlug,
		},
	)
	repo := repositories.NewUserRepository(databases.DB)
	u := usecases.NewUserUseCase(repo)
	handler := deliveries.NewUserHandelr(u, "Users", "User module management", &moduleRoute)

	authRepo := repositories.NewAuthenticationRepository(databases.DB, databases.QueueClient)
	authUseCase := usecases.NewAuthenticationUseCase(authRepo)
	authHandler := deliveries.GetAuthHandlerUsecase(authUseCase)

	e := app.Group("/user", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Post("/", handler.NewUser)
	e.Put("/:id", handler.UpdateUser)
	e.Get("/:id", handler.GetUser)
	e.Delete("/:id", handler.DeleteUser)

	e = app.Group("/users", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Delete("/", handler.DeleteMultitpleUsers)
	e.Get("/", handler.GetAllUsers)
	e.Put("/", handler.RestoreUsers)
}

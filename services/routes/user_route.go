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
	UserSlug     = "/user"
	UsersSlug    = "/users"
	RegisterSlug = "/register"
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
			Name:        "New User",
			Description: "สร้าง User ใหม่",
			Method:      fiber.MethodPost,
			Slug:        UserSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Update User",
			Description: "แก้ไข User",
			Method:      fiber.MethodPut,
			Slug:        UserSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Get User",
			Description: "ดึงข้อมูล User",
			Method:      fiber.MethodGet,
			Slug:        UserSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Delete User",
			Description: "ลบข้อมูล User",
			Method:      fiber.MethodDelete,
			Slug:        UserSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Delete User",
			Description: "ลบข้อมูล User แบบ Mutiple",
			Method:      fiber.MethodDelete,
			Slug:        UsersSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Get All User",
			Description: "ดึงข้อมูล User ทั้งหมด",
			Method:      fiber.MethodGet,
			Slug:        UsersSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Restore User",
			Description: "Restore User จากถังขยะไป inbox",
			Method:      fiber.MethodPut,
			Slug:        UsersSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Register User",
			Description: "Register New User From Frontend",
			Method:      fiber.MethodPost,
			Slug:        RegisterSlug,
		},
	)
	repo := repositories.NewUserRepository(databases.DB)
	u := usecases.NewUserUseCase(repo)
	handler := deliveries.NewUserHandelr(u, &moduleRoute)

	e := app.Group("/register")
	e.Post("/", handler.RegisterUser)

	e = app.Group("/user", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Post("/", handler.NewUser)
	e.Put("/:id", handler.UpdateUser)
	e.Get("/:id", handler.GetUser)
	e.Delete("/:id", handler.DeleteUser)

	e = app.Group("/users", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Delete("/", handler.DeleteMultitpleUsers)
	e.Get("/", handler.GetAllUsers)
	e.Put("/", handler.RestoreUsers)
}

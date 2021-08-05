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
			Name:        "Add New Role",
			Description: "สร้าง Role ใหม่",
			Method:      fiber.MethodPost,
			Slug:        RoleSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Get Role",
			Description: "ดึงข้อมูล Role มาแสดง",
			Method:      fiber.MethodGet,
			Slug:        RoleSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Update Role",
			Description: "แก้ไข Role",
			Method:      fiber.MethodPut,
			Slug:        RoleSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Get All Role",
			Description: "ดึงข้อมูล Role ทั้งหมด",
			Method:      fiber.MethodGet,
			Slug:        RolesSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Delete Role",
			Description: "ลบข้อมูล Role แบบ Multiple",
			Method:      fiber.MethodDelete,
			Slug:        RolesSlug,
		},
		models.ModuleMethodSlug{
			Name:        "Restore Role",
			Description: "ย้ายข้อมูล Role จาก ถังขยะไป inbox แบบ Multitple",
			Method:      fiber.MethodPut,
			Slug:        RestoreRolesSlug,
		},
	)
	repo := repositories.NewRoleRepository(databases.DB)
	u := usecases.NewRoleUseCase(repo)
	handler := deliveries.NewRoleHandelr(u, &moduleRoute)

	e := app.Group("/role", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Post("/", handler.NewRole)
	e.Get("/:id", handler.GetRole)
	e.Put("/:id", handler.UpdateRole)

	e = app.Group("/roles", authHandler.AuthorizationRequired(), authHandler.IsRevokeToken, authHandler.CheckRoutingPermission)
	e.Get("/", handler.GetAllRoles)
	e.Delete("/", handler.DeleteRoles)
	e.Put("/restore", handler.RestoreRoles)
}

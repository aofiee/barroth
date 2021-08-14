package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/gofiber/fiber/v2"
)

func InitAllRoutes() *fiber.App {
	i := NewInstallationRoutes(barroth_config.ENV)
	app := i.Setup()
	i.Install(app)

	role := NewRoleRoutes(barroth_config.ENV)
	role.Install(app)

	user := NewUserRoutes(barroth_config.ENV)
	user.Install(app)

	auth := NewAuthenticationRoutes(barroth_config.ENV)
	auth.Install(app)

	module := NewModuleRoutes(barroth_config.ENV)
	module.Install(app)

	permissions := NewPermissionRoutes(barroth_config.ENV)
	permissions.Install(app)

	forgotpassword := NewForgotPasswordRoutes(barroth_config.ENV)
	forgotpassword.Install(app)
	return app
}

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
	forgotPasswordRoute struct {
		config barroth_config.Config
	}
)

func NewForgotPasswordRoutes(config barroth_config.Config) *forgotPasswordRoute {
	return &forgotPasswordRoute{
		config: config,
	}
}
func (f *forgotPasswordRoute) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Name:        "User Reset Password",
			Description: "User reset password by mail gun APIs",
			Method:      fiber.MethodPut,
			Slug:        "/reset_password",
		},
	)
	repo := repositories.NewForgotPasswordRepository(databases.DB, databases.ResetPasswordQueueClient)
	u := usecases.NewForgotPasswordUseCase(repo)
	handler := deliveries.NewForgotPasswordHandler(u, &moduleRoute)
	e := app.Group("/reset_password")
	e.Put("/:id", handler.ResetPassword)
	e.Get("/:id", handler.ResetPasswordForm)
	e.Post("/:id", handler.ResetPasswordFormExec)
}

package routes

import (
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/deliveries"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

type (
	installationRoutes struct {
		config barroth_config.Config
	}
)

var (
	authRepo    domains.AuthenticationRepository
	authUseCase domains.AuthenticationUseCase
	authHandler *deliveries.AuthenticationHandler
)

func NewInstallationRoutes(config barroth_config.Config) *installationRoutes {
	return &installationRoutes{
		config: config,
	}
}

func (i *installationRoutes) Setup() *fiber.App {
	app := fiber.New()
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format:     "${pid} ${status} - ${method} ${path}\n",
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Asia/Bangkok",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: i.config.AllowOrigins,
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Live!")
	})
	return app
}

func (i *installationRoutes) Install(app *fiber.App) {
	var moduleRoute []models.ModuleMethodSlug
	moduleRoute = append(moduleRoute,
		models.ModuleMethodSlug{
			Name:        "Install",
			Description: "Install Administrator",
			Method:      fiber.MethodGet,
			Slug:        "/install",
		},
	)
	sysRepo := repositories.NewSystemRepository(databases.DB)
	sysUseCase := usecases.NewSystemUseCase(sysRepo)
	installHandler := deliveries.NewSystemHandelr(sysUseCase, &moduleRoute)

	authRepo = repositories.NewAuthenticationRepository(databases.DB, databases.QueueClient)
	authUseCase = usecases.NewAuthenticationUseCase(authRepo)
	authHandler = deliveries.GetAuthHandlerUsecase(authUseCase)

	e := app.Group("/install")
	e.Get("/", installHandler.SystemInstallation)
}

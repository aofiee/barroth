package main

import (
	"log"
	"os"
	"time"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dns, err := setupDNSDatabaseConnection("../")
	if err != nil {
		log.Println(err)
	}
	dial := mysql.Open(dns)
	err = createDatabaseConnection(dial)
	if err != nil {
		log.Println(err)
	}
	/// Install Routing
	r := routes.NewInstallationRoutes(barroth_config.ENV)
	app := r.Setup()
	r.Install(app)
	err = app.Listen(":" + barroth_config.ENV.AppPort)
	if err != nil {
		panic(err)
	}
}

func setupDNSDatabaseConnection(env string) (string, error) {
	var err error
	/// Load Configuration
	barroth_config.ENV, err = barroth_config.LoadConfig(env)
	if err != nil {
		return "", err
	}
	/// Database Connection
	dns := databases.NewConfig(barroth_config.ENV).DBConnString()
	return dns, nil
}

func createDatabaseConnection(dial gorm.Dialector) error {
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	databases.DB, err = gorm.Open(dial, &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		return err
	}
	err = databases.DB.AutoMigrate(&models.Users{}, &models.RoleItems{}, &models.UserRoles{}, &models.Modules{}, &models.Permissions{}, &models.System{})
	return err
}

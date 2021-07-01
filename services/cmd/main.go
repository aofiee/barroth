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
	var err error
	/// Load Configuration
	barroth_config.ENV, err = barroth_config.LoadConfig("../")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/// Database Connection
	dns := databases.NewConfig(barroth_config.ENV).DBConnString()
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)
	databases.DB, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Println("error", err.Error())
		log.Fatal(err)
	}
	databases.DB.AutoMigrate(&models.System{})
	/// Install Routing
	r := routes.NewInstallationRoutes(barroth_config.ENV)
	app := r.Setup()
	r.Install(app)
	err = app.Listen(":" + barroth_config.ENV.AppPort)
	if err != nil {
		panic(err)
	}
}

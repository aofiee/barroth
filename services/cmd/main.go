package main

import (
	"log"
	"os"
	"time"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/routes"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	dbDNS, tokenQueueDNS, resetQueueDNS, err := setupDNSDatabaseConnection("../")
	if err != nil {
		log.Println(err)
	}
	dial := mysql.Open(dbDNS)
	err = createDatabaseConnection(dial)
	if err != nil {
		log.Println(err)
	}
	createTokenQueueConnection(tokenQueueDNS, barroth_config.ENV.TokenRdPassword)
	createResetPasswordQueueConnection(resetQueueDNS, barroth_config.ENV.ResetpasswordRdPassword)
	log.Println("tokenQueueDNS, resetQueueDNS", tokenQueueDNS, resetQueueDNS)
	/// Install Routing
	app := routes.InitAllRoutes()
	///
	err = app.Listen(":" + barroth_config.ENV.AppPort)
	if err != nil {
		panic(err)
	}
}
func createTokenQueueConnection(dns, password string) {
	databases.TokenQueueClient = redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: password,
	})
}
func createResetPasswordQueueConnection(dns, password string) {
	databases.ResetPasswordQueueClient = redis.NewClient(&redis.Options{
		Addr:     dns,
		Password: password,
	})
}
func setupDNSDatabaseConnection(env string) (string, string, string, error) {
	var err error
	/// Load Configuration
	barroth_config.ENV, err = barroth_config.LoadConfig(env)
	if err != nil {
		return "", "", "", err
	}
	/// Database Connection
	databaseDNS := databases.NewConfig(barroth_config.ENV).DBConnString()
	redisTokenDNS := databases.NewConfig(barroth_config.ENV).TokenRedisConnString()
	redisResetPasswordDNS := databases.NewConfig(barroth_config.ENV).ResetPasswordRedisConnString()
	return databaseDNS, redisTokenDNS, redisResetPasswordDNS, nil
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

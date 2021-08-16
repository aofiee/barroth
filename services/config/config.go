package config

import (
	"github.com/spf13/viper"
)

var (
	ENV Config
)

type (
	Config struct {
		AppName                 string `mapstructure:"APP_NAME"`
		AllowOrigins            string `mapstructure:"ALLOW_ORIGINS"`
		DbHost                  string `mapstructure:"DB_HOST"`
		DbPort                  string `mapstructure:"DB_PORT"`
		DbUser                  string `mapstructure:"DB_USER"`
		DbPassword              string `mapstructure:"DB_PASSWORD"`
		DbName                  string `mapstructure:"DB_NAME"`
		RbUser                  string `mapstructure:"RB_USER"`
		RbPassword              string `mapstructure:"RB_PASSWORD"`
		RbHost                  string `mapstructure:"RB_HOST"`
		RbPort                  string `mapstructure:"RB_PORT"`
		TokenRdPassword         string `mapstructure:"TOKEN_REDIS_PASSWORD"`
		TokenRdHost             string `mapstructure:"TOKEN_REDIS_HOST"`
		TokenRdPort             string `mapstructure:"TOKEN_REDIS_PORT"`
		ResetpasswordRdPassword string `mapstructure:"RESETPASSWORD_REDIS_PASSWORD"`
		ResetpasswordRdHost     string `mapstructure:"RESETPASSWORD_REDIS_HOST"`
		ResetpasswordRdPort     string `mapstructure:"RESETPASSWORD_REDIS_PORT"`
		AccessKey               string `mapstructure:"ACCESS_TOKEN_SECRET"`
		RefreshKey              string `mapstructure:"REFRESH_TOKEN_SECRET"`
		AppPort                 string `mapstructure:"APP_PORT"`
		SiteURL                 string `mapstructure:"SITE_URL"`
		EmailAdministrator      string `mapstructure:"EMAIL_ADMINISTRATOR"`
		PasswordAdministrator   string `mapstructure:"PASSWORD_ADMINISTRATOR"`
		TelephoneAdministrator  string `mapstructure:"TELEPHONE_ADMINISTRATOR"`
		MyDomain                string `mapstructure:"MYDOMAIN"`
		MailGunApiKey           string `mapstructure:"MAIL_GUN_PRIVATE_API_KEY"`
	}
)

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("example-app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}

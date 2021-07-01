package main

import (
	"testing"

	barroth_config "github.com/aofiee/barroth/config"
	utils "github.com/gofiber/fiber/v2/utils"
)

func TestMain(t *testing.T) {
	t.Run("TEST_LOAD_ENV", func(t *testing.T) {
		var err error
		barroth_config.ENV, err = barroth_config.LoadConfig("./")
		if err != nil {
			utils.AssertEqual(t, `Config File "app" Not Found in "[/Users/aofiee/Documents/Projects/CleanArchitecture/services/cmd]"`, err.Error(), "barroth_config.ENV")
		}
		barroth_config.ENV, err = barroth_config.LoadConfig("../")
		utils.AssertEqual(t, "Diablos", barroth_config.ENV.AppName, "barroth_config.ENV")
	})
}

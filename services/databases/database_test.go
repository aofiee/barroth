package databases

import (
	"reflect"
	"testing"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("TEST_DATABASE_NEW_CONFIG", func(t *testing.T) {
		var err error
		barroth_config.ENV, err = barroth_config.LoadConfig("../")
		if err != nil {
			assert.NotEqual(t, nil, err, "db config")
		}
		expectedDB := barroth_config.ENV.DbUser + ":" + barroth_config.ENV.DbPassword + "@tcp(database:3306)/" + barroth_config.ENV.DbName + "?charset=utf8mb4&parseTime=True&loc=Local"

		expectedTokenQueue := barroth_config.ENV.TokenRdHost + ":" + barroth_config.ENV.TokenRdPort

		expectedResetPasswordQueue := barroth_config.ENV.ResetpasswordRdHost + ":" + barroth_config.ENV.ResetpasswordRdPort

		db := NewConfig(barroth_config.ENV)
		assert.Equal(t, "*databases.DBConfig", reflect.TypeOf(db).String(), "db config")
		dns := db.DBConnString()
		assert.Equal(t, expectedDB, dns, "dns database")

		queueDNS := db.TokenRedisConnString()
		assert.Equal(t, expectedTokenQueue, queueDNS, "dns queue")

		queueDNS = db.ResetPasswordRedisConnString()
		assert.Equal(t, expectedResetPasswordQueue, queueDNS, "dns queue")
	})
}

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
		db := NewConfig(barroth_config.ENV)
		assert.Equal(t, "*databases.DBConfig", reflect.TypeOf(db).String(), "db config")
		dns := db.DBConnString()
		assert.NotEqual(t, "", dns, "dns database")
	})
}

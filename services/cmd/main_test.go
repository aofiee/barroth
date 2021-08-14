package main

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
)

func SetupMock(t *testing.T) {
	var err error
	barroth_config.ENV, err = barroth_config.LoadConfig("../")
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
	db, mock, err = sqlmock.New()
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
}
func TestSetupDatabase(t *testing.T) {
	SetupMock(t)
	t.Run("TEST_LOAD_FAIL_ENV", func(t *testing.T) {
		dbDNS, queueDNS, err := setupDNSDatabaseConnection("./")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
			assert.Equal(t, "", dbDNS, constants.ERR_DNS_CONNECTION_EMPTY)
			assert.Equal(t, "", queueDNS, constants.ERR_DNS_CONNECTION_EMPTY)
		}
	})
	t.Run("TEST_LOAD_SUCCESS_ENV", func(t *testing.T) {
		dbDNS, queueDNS, err := setupDNSDatabaseConnection("./")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
			assert.Equal(t, "", dbDNS, constants.ERR_DNS_CONNECTION_EMPTY)
			assert.Equal(t, "", queueDNS, constants.ERR_DNS_CONNECTION_EMPTY)
		}
		assert.NotEqual(t, "", dbDNS, constants.ERR_DNS_CONNECTION_EMPTY)
		assert.NotEqual(t, "", queueDNS, constants.ERR_DNS_CONNECTION_EMPTY)
	})
	t.Run("DATABASE_TEST", func(t *testing.T) {
		dbDNS, queueDNS, err := setupDNSDatabaseConnection("./")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
			assert.Equal(t, "", dbDNS, constants.ERR_DNS_CONNECTION_EMPTY)
			assert.Equal(t, "", queueDNS, constants.ERR_DNS_CONNECTION_EMPTY)
		}
		dial := mysql.New(mysql.Config{
			DSN:        "sqlmock_db_0",
			DriverName: "mysql",
			Conn:       db,
		})
		rows := mock.NewRows([]string{"VERSION()"}).
			AddRow("5.7.34")
		assert.Equal(t, "*sqlmock.Rows", reflect.TypeOf(rows).String(), "new row")
		mock.ExpectQuery("SELECT VERSION()").
			WillReturnRows(rows)
		err = createDatabaseConnection(dial)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		createQueueConnection(queueDNS, barroth_config.ENV.TokenRdPassword)

	})
}

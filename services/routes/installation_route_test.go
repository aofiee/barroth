package routes

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db    *sql.DB
	smock sqlmock.Sqlmock
)

func SetupMock(t *testing.T) {
	var err error
	barroth_config.ENV, err = barroth_config.LoadConfig("../")
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
	db, smock, err = sqlmock.New()
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
	dial := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       db,
	})
	rows := smock.NewRows([]string{"VERSION()"}).
		AddRow("5.7.34")
	assert.Equal(t, "*sqlmock.Rows", reflect.TypeOf(rows).String(), "new row")
	smock.ExpectQuery("SELECT VERSION()").
		WillReturnRows(rows)
	databases.DB, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
}
func TestInstallation(t *testing.T) {
	SetupMock(t)
	ins := NewInstallationRoutes(barroth_config.ENV)
	assert.Equal(t, "*routes.installationRoutes", reflect.TypeOf(ins).String(), "new installation")
	app := ins.Setup()
	ins.Install(app)
}

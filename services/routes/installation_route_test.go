package routes

import (
	"database/sql"
	"net/http/httptest"
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
	dial := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       db,
	})
	rows := mock.NewRows([]string{"VERSION()"}).
		AddRow("5.7.34")
	mock.ExpectQuery("SELECT VERSION()").
		WillReturnRows(rows)
	databases.DB, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
}
func TestInstallation(t *testing.T) {
	ins := NewInstallationRoutes(barroth_config.ENV)
	assert.Equal(t, "*routes.installationRoutes", reflect.TypeOf(ins).String(), "new installation")
	app := ins.Setup()

	app.Get("/", nil)
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
	assert.Equal(t, 200, resp.StatusCode, "response body")

	assert.Equal(t, "*fiber.App", reflect.TypeOf(app).String(), "new installation")
	ins.Install(app)
}

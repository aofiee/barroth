package deliveries

import (
	"database/sql"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
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
	assert.Equal(t, "*sqlmock.Rows", reflect.TypeOf(rows).String(), "new row")
	mock.ExpectQuery("SELECT VERSION()").
		WillReturnRows(rows)
	databases.DB, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
}
func TestNewSystemHandelr(t *testing.T) {
	SetupMock(t)
	sysRepo := repositories.NewSystemRepository(databases.DB)
	sysUseCase := usecases.NewSystemUseCase(sysRepo)
	sysHandler := NewSystemHandelr(sysUseCase)
	assert.Equal(t, "*deliveries.systemHandler", reflect.TypeOf(sysHandler).String(), "db config")

	t.Run("TEST_SOFTWARE_IS_INSTALLED", func(t *testing.T) {
		sysRows := mock.NewRows([]string{"id", "app_name", "site_url", "is_install"}).
			AddRow(1, "MyApplication", "http://localhost:8181", 0)
		assert.Equal(t, "*sqlmock.Rows", reflect.TypeOf(sysRows).String(), "new row")
		mock.ExpectQuery("^SELECT (.*)").
			WillReturnRows(sysRows)
		app := fiber.New()
		app.Post("/", sysHandler.SystemInstallation)
		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		assert.Equal(t, 200, resp.StatusCode, "response body")
	})
	t.Run("TEST_SOFTWARE_IS_NOT_INSTALLED", func(t *testing.T) {
		app := fiber.New()
		app.Post("/", sysHandler.SystemInstallation)
		req := httptest.NewRequest("POST", "/", nil)
		req.Header.Set("Content-Type", "application/json")
		resp, err := app.Test(req)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		assert.Equal(t, 400, resp.StatusCode, "response body")
	})
}

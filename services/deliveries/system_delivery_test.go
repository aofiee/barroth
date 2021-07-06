package deliveries

import (
	"database/sql"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
func SystemMockSetup(t *testing.T) (mockUseCase *mocks.SystemUseCase, sysHandler *systemHandler) {
	SetupMock(t)
	var mockSystem models.System
	err := faker.FakeData(&mockSystem)
	assert.NoError(t, err)
	mockUseCase = new(mocks.SystemUseCase)
	sysHandler = NewSystemHandelr(mockUseCase, "expect1", "expect2", "slug")
	return
}
func TestNewSystemHandelrInstallingCompleted(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType("*models.System")).Return(errors.New("hello world"))

	mockUseCase.On("CreateSystem", mock.AnythingOfType("*models.System")).Return(nil)

	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestNewSystemHandelrInstallingFailed(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType("*models.System")).Return(errors.New("error"))
	mockUseCase.On("CreateSystem", mock.AnythingOfType("*models.System")).Return(errors.New("error"))
	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestNewSystemHandelrInstalled(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType("*models.System")).Return(nil)
	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 302, resp.StatusCode, "completed")
}

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

const (
	contentType          = "application/json"
	mockSystemType       = "*models.System"
	mockSystemTypeSlice  = "*[]models.System"
	sliceModuleModelType = "*[]models.Modules"
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
	var mr []models.ModuleMethodSlug
	mr = append(mr, models.ModuleMethodSlug{
		Name:        "Test",
		Description: "Desc Test",
		Method:      fiber.MethodPost,
		Slug:        "/test",
	})
	var mockSystem models.System
	err := faker.FakeData(&mockSystem)
	assert.NoError(t, err)
	mockUseCase = new(mocks.SystemUseCase)
	sysHandler = NewSystemHandelr(mockUseCase, &mr)
	return
}
func TestNewSystemHandelrInstallingCompleted(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType(mockSystemType)).Return(errors.New("hello world"))
	mockUseCase.On("CreateRole", mock.AnythingOfType(mockRoleType)).Return(nil)
	mockUseCase.On("SetExecToAllModules", mock.AnythingOfType(sliceModuleModelType), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil)
	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	mockUseCase.On("CreateSystem", mock.AnythingOfType(mockSystemType)).Return(nil)
	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}

func TestNewSystemHandelrInstallingCreateRoleFail(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType(mockSystemType)).Return(errors.New("hello world"))
	mockUseCase.On("CreateRole", mock.AnythingOfType(mockRoleType)).Return(errors.New("error CreateRole"))

	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}

func TestNewSystemHandelrInstallingSetExecToAllModulesFail(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType(mockSystemType)).Return(errors.New("hello world"))
	mockUseCase.On("CreateRole", mock.AnythingOfType(mockRoleType)).Return(nil)
	mockUseCase.On("SetExecToAllModules", mock.AnythingOfType(sliceModuleModelType), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(errors.New("error SetExecToAllModules"))

	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}

func TestNewSystemHandelrCreateUserFailed(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType(mockSystemType)).Return(errors.New("hello world"))
	mockUseCase.On("CreateRole", mock.AnythingOfType(mockRoleType)).Return(nil)
	mockUseCase.On("SetExecToAllModules", mock.AnythingOfType(sliceModuleModelType), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil)
	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(errors.New("error CreateUser"))

	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestNewSystemHandelrCreateSystemFailed(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType(mockSystemType)).Return(errors.New("hello world"))
	mockUseCase.On("CreateRole", mock.AnythingOfType(mockRoleType)).Return(nil)
	mockUseCase.On("SetExecToAllModules", mock.AnythingOfType(sliceModuleModelType), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil)
	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	mockUseCase.On("CreateSystem", mock.AnythingOfType(mockSystemType)).Return(errors.New("error CreateSystem"))

	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestNewSystemHandelrInstalled(t *testing.T) {
	mockUseCase, sysHandler := SystemMockSetup(t)
	mockUseCase.On("GetFirstSystemInstallation", mock.AnythingOfType(mockSystemType)).Return(nil)
	app := fiber.New()
	app.Get("/install", sysHandler.SystemInstallation)
	req, err := http.NewRequest("GET", "/install", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 302, resp.StatusCode, "completed")
}

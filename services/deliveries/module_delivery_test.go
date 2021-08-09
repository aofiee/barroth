package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	moduleName           = "Update Module"
	moduleDescription    = "Description"
	mockModulesType      = "*models.Modules"
	mockModulesTypeSlice = "*[]models.Modules"
)

func ModulesMockSetup(t *testing.T) (mockUseCase *mocks.ModuleUseCase, handler *moduleHandler) {
	SetupMock(t)
	var mr []models.ModuleMethodSlug
	mr = append(mr, models.ModuleMethodSlug{
		Name:        "Test",
		Description: "Desc Test",
		Method:      fiber.MethodPost,
		Slug:        "/test",
	})
	var mockModules models.Modules
	err := faker.FakeData(&mockModules)
	assert.NoError(t, err)
	mockUseCase = new(mocks.ModuleUseCase)
	handler = NewModuleHandler(mockUseCase, &mr)
	return
}
func TestGetAllModuleValidateFail(t *testing.T) {
	_, handler := ModulesMockSetup(t)
	app := fiber.New()
	app.Get("/modules", handler.GetAllModules)
	req, err := http.NewRequest("GET", "/modules?sort=abab", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestGetAllModuleFail(t *testing.T) {
	mockUseCase, handler := ModulesMockSetup(t)
	mockUseCase.On("GetAllModules", mock.AnythingOfType(mockModulesTypeSlice), "all", "asc", "id", "0", "10", "inbox").Return(errors.New("error"))
	app := fiber.New()
	app.Get("/modules", handler.GetAllModules)
	req, err := http.NewRequest("GET", "/modules", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestGetAllModuleSuccess(t *testing.T) {
	mockUseCase, handler := ModulesMockSetup(t)
	mockUseCase.On("GetAllModules", mock.AnythingOfType(mockModulesTypeSlice), "all", "asc", "id", "0", "10", "inbox").Return(nil)
	app := fiber.New()
	app.Get("/modules", handler.GetAllModules)
	req, err := http.NewRequest("GET", "/modules", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestUpdateModuleJSONNotSend(t *testing.T) {
	_, handler := ModulesMockSetup(t)
	app := fiber.New()
	app.Put("/modules/:id", handler.UpdateModule)
	req, err := http.NewRequest("PUT", "/modules/2", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestUpdateModuleValidateFail(t *testing.T) {
	_, handler := ModulesMockSetup(t)
	params := paramsModules{
		Name:        moduleName,
		Description: moduleDescription,
		Method:      "Hello",
		ModuleSlug:  "/modules",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/modules/:id", handler.UpdateModule)
	req, err := http.NewRequest("PUT", "/modules/2", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestUpdateModuleQueryStringFail(t *testing.T) {
	_, handler := ModulesMockSetup(t)
	params := paramsModules{
		Name:        moduleName,
		Description: moduleDescription,
		Method:      fiber.MethodPut,
		ModuleSlug:  "/modules",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/modules/:id", handler.UpdateModule)
	req, err := http.NewRequest("PUT", "/modules/hello", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestUpdateModuleUpdateModuleFail(t *testing.T) {
	mockUseCase, handler := ModulesMockSetup(t)
	mockUseCase.On("UpdateModule", mock.AnythingOfType(mockModulesType), mock.AnythingOfType("uint")).Return(errors.New("error UpdateModule"))
	params := paramsModules{
		Name:        moduleName,
		Description: moduleDescription,
		Method:      fiber.MethodPut,
		ModuleSlug:  "/modules",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/modules/:id", handler.UpdateModule)
	req, err := http.NewRequest("PUT", "/modules/2", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestUpdateModuleUpdateModuleSuccess(t *testing.T) {
	mockUseCase, handler := ModulesMockSetup(t)
	mockUseCase.On("UpdateModule", mock.AnythingOfType(mockModulesType), mock.AnythingOfType("uint")).Return(nil)
	params := paramsModules{
		Name:        moduleName,
		Description: moduleDescription,
		Method:      fiber.MethodPut,
		ModuleSlug:  "/modules",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/modules/:id", handler.UpdateModule)
	req, err := http.NewRequest("PUT", "/modules/2", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}

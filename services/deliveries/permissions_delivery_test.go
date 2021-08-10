package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	mockPermissionsTypeSlice = "*[]models.Permissions"
)

func PermissionsMockSetup(t *testing.T) (mockUseCase *mocks.PermissionsUseCase, handler *permissionsHandler) {
	SetupMock(t)
	var mr []models.ModuleMethodSlug
	mr = append(mr, models.ModuleMethodSlug{
		Name:        "Test",
		Description: "Desc Test",
		Method:      fiber.MethodPost,
		Slug:        "/test",
	})
	mockUseCase = new(mocks.PermissionsUseCase)
	handler = NewPermissionHandler(mockUseCase, &mr)
	return
}
func TestSetPermissionsJSONNotSend(t *testing.T) {
	_, handler := PermissionsMockSetup(t)
	app := fiber.New()
	app.Put("/permissions/:id", handler.SetPermission)
	req, err := http.NewRequest("PUT", "/permissions/1", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestSetPermissionsJSONValidateFail(t *testing.T) {
	_, handler := PermissionsMockSetup(t)
	app := fiber.New()
	params := paramsSetPermission{}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Put("/permissions/:id", handler.SetPermission)
	req, err := http.NewRequest("PUT", "/permissions/1", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestSetPermissionsQueryStringFail(t *testing.T) {
	_, handler := PermissionsMockSetup(t)
	app := fiber.New()
	var permissions []permission
	permissions = append(permissions, permission{
		ModuleID: 1,
		IsExec:   1,
	})
	params := paramsSetPermission{
		Permissions: permissions,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Put("/permissions/:id", handler.SetPermission)
	req, err := http.NewRequest("PUT", "/permissions/helloworld", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestSetPermissionsSetPermissionsFail(t *testing.T) {
	mockUseCase, handler := PermissionsMockSetup(t)
	app := fiber.New()
	var permissions []permission
	permissions = append(permissions, permission{
		ModuleID: 1,
		IsExec:   1,
	})
	params := paramsSetPermission{
		Permissions: permissions,
	}
	mockUseCase.On("SetPermissions", mock.AnythingOfType(mockPermissionsTypeSlice)).Return(errors.New("error SetPermissions"))
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Put("/permissions/:id", handler.SetPermission)
	req, err := http.NewRequest("PUT", "/permissions/1", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestSetPermissionsSetPermissionsSuccess(t *testing.T) {
	mockUseCase, handler := PermissionsMockSetup(t)
	app := fiber.New()
	var permissions []permission
	permissions = append(permissions, permission{
		ModuleID: 1,
		IsExec:   1,
	})
	params := paramsSetPermission{
		Permissions: permissions,
	}
	mockUseCase.On("SetPermissions", mock.AnythingOfType(mockPermissionsTypeSlice)).Return(nil)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Put("/permissions/:id", handler.SetPermission)
	req, err := http.NewRequest("PUT", "/permissions/1", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}

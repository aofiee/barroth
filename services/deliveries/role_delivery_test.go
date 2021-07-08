package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func RoleMockSetup(t *testing.T) (mockUseCase *mocks.RoleUseCase, handler *roleHandler) {
	SetupMock(t)
	var mockRole models.RoleItems
	err := faker.FakeData(&mockRole)
	assert.NoError(t, err)
	mockUseCase = new(mocks.RoleUseCase)
	handler = NewRoleHandelr(mockUseCase, "expect1", "expect2", "/role")
	return
}
func TestNewHandlerSuccess(t *testing.T) {
	params := models.RoleItems{
		Name:        "TestRole",
		Description: "Lorem Test",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("CreateRole", mock.AnythingOfType("*models.RoleItems")).Return(nil)

	app := fiber.New()
	app.Post("/role", handler.NewRole)
	req, err := http.NewRequest("POST", "/role", payload)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
	//case empty json
	app.Post("/role", handler.NewRole)
	req, err = http.NewRequest("POST", "/role", nil)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
	//case validator
	params.Name = ""
	params.Description = ""
	data, _ = json.Marshal(&params)
	payload = bytes.NewReader(data)
	app.Post("/role", handler.NewRole)
	req, err = http.NewRequest("POST", "/role", payload)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestNewHandlerFail(t *testing.T) {
	params := models.RoleItems{
		Name:        "TestRole",
		Description: "Lorem Test",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("CreateRole", mock.AnythingOfType("*models.RoleItems")).Return(errors.New("error"))

	app := fiber.New()
	app.Post("/role", handler.NewRole)
	req, err := http.NewRequest("POST", "/role", payload)
	req.Header.Set("Content-Type", "application/json")
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")

}
func TestGetAllRolesSuccess(t *testing.T) {
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("GetAllRoles", mock.AnythingOfType("*[]models.RoleItems"), "all", "asc", "id", "0", "10", "inbox").Return(nil)
	app := fiber.New()
	app.Get("/roles", handler.GetAllRoles)
	req, err := http.NewRequest("GET", "/roles", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestGetAllRolesFail(t *testing.T) {
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("GetAllRoles", mock.AnythingOfType("*[]models.RoleItems"), "all", "asc", "id", "0", "10", "inbox").Return(errors.New("error"))
	app := fiber.New()
	app.Get("/roles", handler.GetAllRoles)
	req, err := http.NewRequest("GET", "/roles", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")

	mockUseCase.On("GetAllRoles", mock.AnythingOfType("*[]models.RoleItems"), "all", "asc", "id", "0", "10", "inbox").Return(errors.New("error"))
	app.Get("/roles", handler.GetAllRoles)
	req, err = http.NewRequest("GET", "/roles?sort=abab", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestDeleteRoleSuccess(t *testing.T) {
	type paramDeleteRoles struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
	var param paramDeleteRoles
	param.RoleID = []int{1, 2, 3}
	data, _ := json.Marshal(&param)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("DeleteRoles", "inbox", mock.AnythingOfType("[]int")).Return(int64(0), nil)
	app := fiber.New()
	app.Delete("/roles", handler.DeleteRoles)
	req, err := http.NewRequest("DELETE", "/roles", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestDeleteRoleFail(t *testing.T) {
	type paramDeleteRoles struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
	var param paramDeleteRoles
	param.RoleID = []int{1, 2, 3}
	data, _ := json.Marshal(&param)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("DeleteRoles", "inbox", mock.AnythingOfType("[]int")).Return(int64(0), errors.New("error"))
	app := fiber.New()
	app.Delete("/roles", handler.DeleteRoles)
	req, err := http.NewRequest("DELETE", "/roles", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}

func TestDeleteRoleValidateFail(t *testing.T) {
	type paramDeleteRoles struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
	var param paramDeleteRoles
	param.RoleID = nil
	data, _ := json.Marshal(&param)
	payload := bytes.NewReader(data)
	_, handler := RoleMockSetup(t)
	app := fiber.New()
	app.Delete("/roles", handler.DeleteRoles)
	req, err := http.NewRequest("DELETE", "/roles", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestDeleteRoleParseJsonFail(t *testing.T) {
	payload := strings.NewReader("{NAME}")
	_, handler := RoleMockSetup(t)
	app := fiber.New()
	app.Delete("/roles", handler.DeleteRoles)
	req, err := http.NewRequest("DELETE", "/roles", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRestoreRolesSuccess(t *testing.T) {
	type paramDeleteRoles struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
	var param paramDeleteRoles
	param.RoleID = []int{1, 2, 3}
	data, _ := json.Marshal(&param)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("RestoreRoles", mock.AnythingOfType("[]int")).Return(int64(3), nil)
	app := fiber.New()
	app.Put("/roles/restore", handler.RestoreRoles)
	req, err := http.NewRequest("PUT", "/roles/restore", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestRestoreRolesFail(t *testing.T) {
	type paramDeleteRoles struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
	var param paramDeleteRoles
	param.RoleID = []int{1, 2, 3}
	data, _ := json.Marshal(&param)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("RestoreRoles", mock.AnythingOfType("[]int")).Return(int64(0), errors.New("error"))
	app := fiber.New()
	app.Put("/roles/restore", handler.RestoreRoles)
	req, err := http.NewRequest("PUT", "/roles/restore", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRestoreRolesValidateFail(t *testing.T) {
	type paramDeleteRoles struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
	var param paramDeleteRoles
	param.RoleID = nil
	data, _ := json.Marshal(&param)
	payload := bytes.NewReader(data)
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("RestoreRoles", mock.AnythingOfType("[]int")).Return(int64(0), nil)
	app := fiber.New()
	app.Put("/roles/restore", handler.RestoreRoles)
	req, err := http.NewRequest("PUT", "/roles/restore", payload)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}

func TestRestoreRolesJsonFail(t *testing.T) {
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("RestoreRoles", mock.AnythingOfType("[]int")).Return(int64(0), nil)
	app := fiber.New()
	app.Put("/roles/restore", handler.RestoreRoles)
	req, err := http.NewRequest("PUT", "/roles/restore", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestGetRoleSuccess(t *testing.T) {
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("GetRole", mock.AnythingOfType("*models.RoleItems"), mock.Anything).Return(nil)
	app := fiber.New()
	app.Get("/role/:id", handler.GetRole)
	req, err := http.NewRequest("GET", "/role/2", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestGetRoleFail(t *testing.T) {
	mockUseCase, handler := RoleMockSetup(t)
	mockUseCase.On("GetRole", mock.AnythingOfType("*models.RoleItems"), mock.Anything).Return(errors.New("error"))
	app := fiber.New()
	app.Get("/role/:id", handler.GetRole)
	req, err := http.NewRequest("GET", "/role/2", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}

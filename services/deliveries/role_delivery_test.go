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
	assert.Equal(t, 400, resp.StatusCode, "completed")
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

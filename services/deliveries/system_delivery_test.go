package deliveries

import (
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

func TMockSetup(t *testing.T) (mockUseCase *mocks.SystemUseCase, sysHandler *systemHandler) {
	var mockSystem models.System
	err := faker.FakeData(&mockSystem)
	assert.NoError(t, err)
	mockUseCase = new(mocks.SystemUseCase)
	sysHandler = NewSystemHandelr(mockUseCase, "expect1", "expect2")
	return
}
func TestNewSystemHandelrInstallingCompleted(t *testing.T) {
	mockUseCase, sysHandler := TMockSetup(t)
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
	mockUseCase, sysHandler := TMockSetup(t)
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
	mockUseCase, sysHandler := TMockSetup(t)
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

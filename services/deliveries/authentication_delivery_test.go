package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	mockAuthType            = "*models.Users"
	mockAuthTokenDetailType = "*models.TokenDetail"
	mockAuthTypeSlice       = "*[]models.Users"
	authEmail               = "aofiee666@gmail.com"
	authPassword            = "password"
)

func AuthMockSetup(t *testing.T) (mockUseCase *mocks.AuthenticationUseCase, handler *authenticationHandler) {
	SetupMock(t)
	var mockUser models.Users
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUseCase = new(mocks.AuthenticationUseCase)
	handler = NewAuthenHandler(mockUseCase, "expect1", "expect2", "/role")
	return
}

func TestNewAuthenticationHandlerSuccess(t *testing.T) {
	var token models.TokenDetail
	params := paramsLogin{
		Email:    authEmail,
		Password: authPassword,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := AuthMockSetup(t)

	mockUseCase.On("Login", mock.AnythingOfType(mockAuthType), mock.Anything, mock.Anything).Return(nil)

	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(token, nil)

	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)

	mockUseCase.On("GenerateRefreshTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)

	app := fiber.New()
	app.Post("/auth", handler.Login)
	req, err := http.NewRequest("POST", "/auth", payload)

	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestNewAuthenticationHandlerSuccess")
	// assert.Equal(t, nil, token)

	req, err = http.NewRequest("POST", "/auth", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestNewAuthenticationHandlerSuccess Json Error")

	params2 := paramsLogin{
		Email:    "",
		Password: authPassword,
	}
	data2, _ := json.Marshal(&params2)
	payload2 := bytes.NewReader(data2)
	req, err = http.NewRequest("POST", "/auth", payload2)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "TestNewAuthenticationHandlerSuccess Json Error")
}
func TestNewAuthenticationHandlerFail(t *testing.T) {
	params := paramsLogin{
		Email:    authEmail,
		Password: authPassword,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := AuthMockSetup(t)

	mockUseCase.On("Login", mock.AnythingOfType(mockAuthType), mock.Anything, mock.Anything).Return(errors.New(constants.ERR_USERNAME_PASSWORD_INCORRECT))

	app := fiber.New()
	app.Post("/auth", handler.Login)
	req, err := http.NewRequest("POST", "/auth", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode, "TestNewAuthenticationHandlerFail")
}

func TestCreateTokenFail(t *testing.T) {
	var token models.TokenDetail
	params := paramsLogin{
		Email:    authEmail,
		Password: authPassword,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := AuthMockSetup(t)

	mockUseCase.On("Login", mock.AnythingOfType(mockAuthType), mock.Anything, mock.Anything).Return(nil)

	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(token, errors.New("create token error"))

	app := fiber.New()
	app.Post("/auth", handler.Login)
	req, err := http.NewRequest("POST", "/auth", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestNewAuthenticationHandlerSuccess")
}
func TestGenAccessTokenFail(t *testing.T) {
	var token models.TokenDetail
	params := paramsLogin{
		Email:    authEmail,
		Password: authPassword,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := AuthMockSetup(t)

	mockUseCase.On("Login", mock.AnythingOfType(mockAuthType), mock.Anything, mock.Anything).Return(nil)

	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(token, nil)

	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(errors.New("GenerateAccessTokenBy error"))

	app := fiber.New()
	app.Post("/auth", handler.Login)
	req, err := http.NewRequest("POST", "/auth", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestNewAuthenticationHandlerSuccess")
}
func TestGenRefreshTokenFail(t *testing.T) {
	var token models.TokenDetail
	params := paramsLogin{
		Email:    authEmail,
		Password: authPassword,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase, handler := AuthMockSetup(t)

	mockUseCase.On("Login", mock.AnythingOfType(mockAuthType), mock.Anything, mock.Anything).Return(nil)

	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(token, nil)

	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)

	mockUseCase.On("GenerateRefreshTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(errors.New("GenerateRefreshTokenBy error"))

	app := fiber.New()
	app.Post("/auth", handler.Login)
	req, err := http.NewRequest("POST", "/auth", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestNewAuthenticationHandlerSuccess")
}

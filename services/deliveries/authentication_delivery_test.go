package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"
	"time"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	timeLoc                 = "Asia/Bangkok"
	mockAuthType            = "*models.Users"
	mockAuthTokenDetailType = "*models.TokenDetail"
	mockAuthTypeSlice       = "*[]models.Users"
	authEmail               = "aofiee666@gmail.com"
	authPassword            = "password"
	authRoleName            = "Role Name"
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

	mockUseCase.On("SaveToken", mock.Anything, mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)

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
func TestSaveTokenFail(t *testing.T) {
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

	mockUseCase.On("SaveToken", mock.Anything, mock.AnythingOfType(mockAuthTokenDetailType)).Return(errors.New("save token error"))

	app := fiber.New()
	app.Post("/auth", handler.Login)
	req, err := http.NewRequest("POST", "/auth", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode, "TestNewAuthenticationHandlerSuccess")
}
func mockToken() (models.TokenDetail, error) {
	var token models.TokenDetail
	location, _ := time.LoadLocation(timeLoc)
	accessToken := time.Now().In(location).Add(time.Minute * 15).Unix()
	refreshToken := time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	token.AccessTokenExp = accessToken
	token.RefreshTokenExp = refreshToken
	token.AccessUUID = utils.UUIDv4()
	token.RefreshUUID = utils.UUIDv4()

	context := models.TokenContext{
		Email:       authEmail,
		DisplayName: mock.Anything,
	}
	context.Role = authRoleName
	token.Context = context

	//
	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["iss"] = barroth_config.ENV.AppName
	claims["sub"] = utils.UUIDv4()
	claims["exp"] = token.AccessTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["context"] = token.Context
	claims["access_uuid"] = token.AccessUUID
	rs, err := tk.SignedString([]byte(barroth_config.ENV.AccessKey))
	token.Token.AccessToken = rs
	return token, err
}
func TestLogoutSuccess(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	app := fiber.New()

	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)

	app.Delete("/auth/logout", handler.AuthorizationRequired(), handler.Logout)
	req, err := http.NewRequest("DELETE", "/auth/logout", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token.Token.AccessToken)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestLogoutSuccess")

}
func TestLogoutDeleteFail(t *testing.T) {
	token, err := mockToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	//

	mockUseCase, handler := AuthMockSetup(t)
	app := fiber.New()

	mockUseCase.On("DeleteToken", mock.Anything).Return(errors.New("delete error"))

	app.Delete("/auth/logout", handler.AuthorizationRequired(), handler.Logout)
	req, err := http.NewRequest("DELETE", "/auth/logout", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token.Token.AccessToken)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode, "TestLogoutSuccess")

}
func TestLogoutMiddlewareFail(t *testing.T) {
	token, err := mockToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	//

	_, handler := AuthMockSetup(t)
	app := fiber.New()

	app.Delete("/auth/logout", handler.AuthorizationRequired(), handler.Logout)
	req, err := http.NewRequest("DELETE", "/auth/logout", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer ")

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode, "TestLogoutSuccess")

}

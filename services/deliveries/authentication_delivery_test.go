package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"reflect"
	"testing"
	"time"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/golang-jwt/jwt"
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

var UUID = utils.UUIDv4()

func AuthMockSetup(t *testing.T) (mockUseCase *mocks.AuthenticationUseCase, handler *authenticationHandler) {
	SetupMock(t)
	var mr []models.ModuleMethodSlug
	mr = append(mr, models.ModuleMethodSlug{
		Method: fiber.MethodPost,
		Slug:   "/test",
	})
	var mockUser models.Users
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUseCase = new(mocks.AuthenticationUseCase)
	handler = NewAuthenHandler(mockUseCase, "expect1", "expect2", &mr)
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
	//fiber.MIMEApplicationJSON
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
func TestGetAuthHandlerUsecase(t *testing.T) {
	SetupMock(t)
	mockUseCase := new(mocks.AuthenticationUseCase)
	handler := GetAuthHandlerUsecase(mockUseCase)
	assert.Equal(t, "*deliveries.authenticationHandler", reflect.TypeOf(handler).String())
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
func mockAccessToken() (models.TokenDetail, error) {
	var token models.TokenDetail
	location, _ := time.LoadLocation(timeLoc)
	accessToken := time.Now().In(location).Add(time.Minute * 15).Unix()
	refreshToken := time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	token.AccessTokenExp = accessToken
	token.RefreshTokenExp = refreshToken
	token.AccessUUID = UUID
	token.RefreshUUID = UUID

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
	claims["sub"] = UUID
	claims["exp"] = token.AccessTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["access_uuid"] = token.AccessUUID
	rs, err := tk.SignedString([]byte(barroth_config.ENV.AccessKey))
	token.Token.AccessToken = rs
	return token, err
}
func TestLogoutSuccess(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)

	token, err := mockAccessToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	app := fiber.New()

	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)

	app.Delete("/auth/logout", handler.AuthorizationRequired(), handler.Logout)
	req, err := http.NewRequest("DELETE", "/auth/logout", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token.Token.AccessToken)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestLogoutSuccess")

}
func TestLogoutDeleteFail(t *testing.T) {
	token, err := mockAccessToken()
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
	assert.Equal(t, 400, resp.StatusCode, "TestLogoutSuccess")

}
func TestLogoutMiddlewareFail(t *testing.T) {
	_, handler := AuthMockSetup(t)
	token, err := mockAccessToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
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
func mockRefreshToken(key string) (models.TokenDetail, error) {
	var token models.TokenDetail
	location, _ := time.LoadLocation(timeLoc)
	accessToken := time.Now().In(location).Add(time.Minute * 15).Unix()
	refreshToken := time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	token.AccessTokenExp = accessToken
	token.RefreshTokenExp = refreshToken
	token.AccessUUID = UUID
	token.RefreshUUID = UUID

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
	claims["sub"] = UUID

	claims["exp"] = token.RefreshTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["refresh_uuid"] = token.RefreshUUID
	rs, err := tk.SignedString([]byte(key))
	token.Token.RefreshToken = rs
	return token, err
}
func TestRefreshTokenFaileParamFake(t *testing.T) {
	_, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFaileTokenError(t *testing.T) {
	_, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	params := models.RefreshToken{
		Token: "xxxx",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode, "TestRefreshTokenSuccess")
}
func mockRefreshTokenValidSubType(key string) (models.TokenDetail, error) {
	var token models.TokenDetail
	location, _ := time.LoadLocation(timeLoc)
	accessToken := time.Now().In(location).Add(time.Minute * 15).Unix()
	refreshToken := time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	token.AccessTokenExp = accessToken
	token.RefreshTokenExp = refreshToken
	token.AccessUUID = UUID
	token.RefreshUUID = UUID

	context := models.TokenContext{
		Email:       authEmail,
		DisplayName: mock.Anything,
	}
	context.Role = authRoleName
	token.Context = context

	tk := jwt.New(jwt.SigningMethodHS256)
	claims := tk.Claims.(jwt.MapClaims)
	claims["iss"] = barroth_config.ENV.AppName
	claims["sub"] = int64(0)

	claims["exp"] = token.RefreshTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["refresh_uuid"] = token.RefreshUUID
	rs, err := tk.SignedString([]byte(key))
	token.Token.RefreshToken = rs
	return token, err
}
func TestRefreshTokenFaileValidSubType(t *testing.T) {
	_, handler := AuthMockSetup(t)
	token, err := mockRefreshTokenValidSubType(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestRefreshTokenSuccess")
}
func mockRefreshTokenValidRefreshUUIDType(key string) (models.TokenDetail, error) {
	var token models.TokenDetail
	location, _ := time.LoadLocation(timeLoc)
	accessToken := time.Now().In(location).Add(time.Minute * 15).Unix()
	refreshToken := time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	token.AccessTokenExp = accessToken
	token.RefreshTokenExp = refreshToken
	token.AccessUUID = UUID
	token.RefreshUUID = UUID

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
	claims["sub"] = UUID

	claims["exp"] = token.RefreshTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["refresh_uuid"] = int64(0)
	rs, err := tk.SignedString([]byte(key))
	token.Token.RefreshToken = rs
	return token, err
}
func TestRefreshTokenFaileValidRefreshType(t *testing.T) {
	_, handler := AuthMockSetup(t)
	token, err := mockRefreshTokenValidRefreshUUIDType(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFaileOnDeleteError(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(errors.New("error"))

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFaileGetUser(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockAuthType), UUID).Return(errors.New("error"))
	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 404, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFailCreateToken(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	var tk models.TokenDetail
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockAuthType), UUID).Return(nil)
	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(tk, errors.New("error"))

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFailGenToken1(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	var tk models.TokenDetail
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockAuthType), UUID).Return(nil)
	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(tk, nil)
	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(errors.New("error"))

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFailGenToken2(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	var tk models.TokenDetail
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockAuthType), UUID).Return(nil)
	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(tk, nil)
	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)
	mockUseCase.On("GenerateRefreshTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(errors.New("error"))

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenFailSaveToken(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	var tk models.TokenDetail
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockAuthType), UUID).Return(nil)
	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(tk, nil)
	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)
	mockUseCase.On("GenerateRefreshTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)

	mockUseCase.On("SaveToken", mock.Anything, mock.AnythingOfType(mockAuthTokenDetailType)).Return(errors.New("error"))

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 500, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestRefreshTokenSuccess(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)
	token, err := mockRefreshToken(barroth_config.ENV.RefreshKey)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.RefreshToken)
	app := fiber.New()

	var tk models.TokenDetail
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockAuthType), UUID).Return(nil)
	mockUseCase.On("CreateToken", mock.AnythingOfType(mockAuthType)).Return(tk, nil)
	mockUseCase.On("GenerateAccessTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)
	mockUseCase.On("GenerateRefreshTokenBy", mock.AnythingOfType(mockAuthType), mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)
	mockUseCase.On("SaveToken", mock.Anything, mock.AnythingOfType(mockAuthTokenDetailType)).Return(nil)

	params := models.RefreshToken{
		Token: token.Token.RefreshToken,
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Post("/auth/refresh_token", handler.RefreshToken)
	req, err := http.NewRequest("POST", "/auth/refresh_token", payload)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestRefreshTokenSuccess")
}
func TestIsRevokeTokenSuccess(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)

	token, err := mockAccessToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	app := fiber.New()

	mockUseCase.On("GetAccessUUIDFromRedis", mock.AnythingOfType("string")).Return(UUID, nil)
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)

	app.Delete("/auth/logout", handler.AuthorizationRequired(), handler.IsRevokeToken, handler.Logout)
	req, err := http.NewRequest("DELETE", "/auth/logout", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token.Token.AccessToken)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "TestLogoutSuccess")
}
func TestIsRevokeTokenFail(t *testing.T) {
	mockUseCase, handler := AuthMockSetup(t)

	token, err := mockAccessToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	app := fiber.New()

	mockUseCase.On("GetAccessUUIDFromRedis", mock.AnythingOfType("string")).Return(UUID, errors.New("error GetAccessUUIDFromRedis"))
	mockUseCase.On("DeleteToken", mock.AnythingOfType("string")).Return(nil)

	app.Delete("/auth/logout", handler.AuthorizationRequired(), handler.IsRevokeToken, handler.Logout)
	req, err := http.NewRequest("DELETE", "/auth/logout", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token.Token.AccessToken)

	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode, "TestLogoutSuccess")
}

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
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	mockUserType      = "*models.Users"
	mockUserRolesType = "*models.UserRoles"
	mockUserTypeSlice = "*[]models.Users"
	userEmail         = "khomkrid@twinsynergy.co.th"
	userPassword      = "password"
	userFullName      = "Arashi L."
	userTelephone     = "0925905444"
	userRole          = 1
)

func getUser(email, password, name, telephone string, role int) paramsUser {
	return paramsUser{
		Email:     email,
		Password:  password,
		Name:      name,
		Telephone: telephone,
		RoleID:    role,
	}
}
func UserMockSetup(t *testing.T) (mockUseCase *mocks.UserUseCase, handler *userHandler) {
	SetupMock(t)
	var mr []models.ModuleMethodSlug
	mr = append(mr, models.ModuleMethodSlug{
		Name:        "Test",
		Description: "Desc Test",
		Method:      fiber.MethodPost,
		Slug:        "/test",
	})
	var mockUser models.Users
	err := faker.FakeData(&mockUser)
	assert.NoError(t, err)
	mockUseCase = new(mocks.UserUseCase)
	handler = NewUserHandelr(mockUseCase, &mr)
	return
}
func TestRegisterUserHandlerJsonNotSend(t *testing.T) {
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Post("/register", handler.RegisterUser)
	req, err := http.NewRequest("POST", "/register", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRegisterUserHandlerJsonValidateFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := paramsRegister{
		Email:       userEmail,
		UUID:        utils.UUIDv4(),
		Provider:    "-",
		DisplayName: userFullName,
		PhotoURL:    "-",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	app.Post("/register", handler.RegisterUser)
	req, err := http.NewRequest("POST", "/register", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestRegisterUserHandlerSetUserRoleFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := paramsRegister{
		Email:       userEmail,
		UUID:        utils.UUIDv4(),
		Provider:    "facebook.com",
		DisplayName: userFullName,
		PhotoURL:    "-",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	mockUseCase.On("SetUserRole", mock.AnythingOfType(mockUserRolesType), uint(0)).Return(errors.New("error to set user role"))
	app.Post("/register", handler.RegisterUser)
	req, err := http.NewRequest("POST", "/register", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRegisterUserHandlerFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := paramsRegister{
		Email:       userEmail,
		UUID:        utils.UUIDv4(),
		Provider:    "google.com",
		DisplayName: userFullName,
		PhotoURL:    "-",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(errors.New("error TestNewUserHandlerFail"))
	app.Post("/register", handler.RegisterUser)
	req, err := http.NewRequest("POST", "/register", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRegisterUserHandlerSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := paramsRegister{
		Email:       userEmail,
		UUID:        utils.UUIDv4(),
		Provider:    "apple.com",
		DisplayName: userFullName,
		PhotoURL:    "-",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	mockUseCase.On("SetUserRole", mock.AnythingOfType(mockUserRolesType), uint(0)).Return(nil)
	app.Post("/register", handler.RegisterUser)
	req, err := http.NewRequest("POST", "/register", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestNewUserHandlerJsonNotSend(t *testing.T) {
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestNewUserHandlerJsonValidateFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := getUser(userEmail, "", userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestNewUserHandlerSetUserRoleFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	//fail validate
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	mockUseCase.On("SetUserRole", mock.AnythingOfType(mockUserRolesType), uint(0)).Return(errors.New("error to set user role"))
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestNewUserHandlerSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	//fail validate
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(nil)
	mockUseCase.On("SetUserRole", mock.AnythingOfType(mockUserRolesType), uint(0)).Return(nil)
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestNewUserHandlerFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)

	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	mockUseCase.On("CreateUser", mock.AnythingOfType(mockUserType)).Return(errors.New("error TestNewUserHandlerFail"))
	app.Post("/user", handler.NewUser)
	req, err := http.NewRequest("POST", "/user", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestUpdateUserJSONNotSend(t *testing.T) {
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Put("/user/:id", handler.UpdateUser)
	req, err := http.NewRequest("PUT", "/user/"+UUID, nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestUpdateUserJsonValidateFail(t *testing.T) {
	_, handler := UserMockSetup(t)
	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	params.Email = ""
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app.Put("/user/:id", handler.UpdateUser)
	req, err := http.NewRequest("PUT", "/user/"+UUID, payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestUpdateUserFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase.On("UpdateUser", mock.AnythingOfType(mockUserType), mock.Anything).Return(errors.New("error TestNewUserHandlerFail"))
	app.Put("/user/:id", handler.UpdateUser)
	req, err := http.NewRequest("PUT", "/user/"+UUID, payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestUpdateUserGetUserFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase.On("UpdateUser", mock.AnythingOfType(mockUserType), mock.Anything).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockUserType), UUID).Return(errors.New("error user not found"))
	app.Put("/user/:id", handler.UpdateUser)
	req, err := http.NewRequest("PUT", "/user/"+UUID, payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestUpdateUserSetUserRoleFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase.On("UpdateUser", mock.AnythingOfType(mockUserType), mock.Anything).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockUserType), UUID).Return(nil)
	mockUseCase.On("SetUserRole", mock.AnythingOfType(mockUserRolesType), uint(0)).Return(errors.New("error set user role"))
	app.Put("/user/:id", handler.UpdateUser)
	req, err := http.NewRequest("PUT", "/user/"+UUID, payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestUpdateUserSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	app := fiber.New()
	params := getUser(userEmail, userPassword, userFullName, userTelephone, userRole)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	mockUseCase.On("UpdateUser", mock.AnythingOfType(mockUserType), mock.Anything).Return(nil)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockUserType), UUID).Return(nil)
	mockUseCase.On("SetUserRole", mock.AnythingOfType(mockUserRolesType), uint(0)).Return(nil)
	app.Put("/user/:id", handler.UpdateUser)
	req, err := http.NewRequest("PUT", "/user/"+UUID, payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestGetUserFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockUserType), UUID).Return(errors.New("error TestGetUserFail"))
	app := fiber.New()
	app.Get("/user/:id", handler.GetUser)
	req, err := http.NewRequest("GET", "/user/"+UUID, nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestGetUserSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockUserType), UUID).Return(nil)
	app := fiber.New()
	app.Get("/user/:id", handler.GetUser)
	req, err := http.NewRequest("GET", "/user/"+UUID, nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestGetUserMeWithToken(t *testing.T) {
	token, err := mockAccessToken()
	assert.NoError(t, err)
	assert.NotEqual(t, nil, token.Token.AccessToken)
	mockUseCase, handler := UserMockSetup(t)
	_, handlerAuth := AuthMockSetup(t)
	mockUseCase.On("GetUser", mock.AnythingOfType(mockUserType), UUID).Return(nil)
	app := fiber.New()
	app.Get("/user/:id", handlerAuth.AuthorizationRequired(), handler.GetUser)
	req, err := http.NewRequest("GET", "/user/me", nil)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	req.Header.Set(fiber.HeaderAuthorization, "Bearer "+token.Token.AccessToken)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestDeleteUserFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	uuids := []string{
		UUID,
	}
	mockUseCase.On("DeleteUsers", "inbox", uuids).Return(int64(0), errors.New("error TestDeleteUserFail"))
	app := fiber.New()
	app.Delete("/user/:id", handler.DeleteUser)
	req, err := http.NewRequest("DELETE", "/user/"+UUID, nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestDeleteUserSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	uuids := []string{
		UUID,
	}
	mockUseCase.On("DeleteUsers", "inbox", uuids).Return(int64(1), nil)
	app := fiber.New()
	app.Delete("/user/:id", handler.DeleteUser)
	req, err := http.NewRequest("DELETE", "/user/"+UUID, nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestDeleteMultitpleUsersJSONNotSend(t *testing.T) {
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Delete("/users", handler.DeleteMultitpleUsers)
	req, err := http.NewRequest("DELETE", "/users", nil)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestDeleteMultitpleUsersValidateError(t *testing.T) {
	var uuid paramUUID
	data, _ := json.Marshal(&uuid)
	payload := bytes.NewReader(data)
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Delete("/users", handler.DeleteMultitpleUsers)
	req, err := http.NewRequest("DELETE", "/users", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestDeleteMultitpleUsersDeleteFail(t *testing.T) {
	var uuid paramUUID
	uuid.UsersID = []string{
		UUID,
	}
	data, _ := json.Marshal(&uuid)
	payload := bytes.NewReader(data)
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("DeleteUsers", "inbox", uuid.UsersID).Return(int64(0), errors.New("error TestDeleteMultitpleUsersDeleteFail"))
	app := fiber.New()
	app.Delete("/users", handler.DeleteMultitpleUsers)
	req, err := http.NewRequest("DELETE", "/users", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestDeleteMultitpleUsersDeleteSuccess(t *testing.T) {
	var uuid paramUUID
	uuid.UsersID = []string{
		UUID,
	}
	data, _ := json.Marshal(&uuid)
	payload := bytes.NewReader(data)
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("DeleteUsers", "inbox", uuid.UsersID).Return(int64(1), nil)
	app := fiber.New()
	app.Delete("/users", handler.DeleteMultitpleUsers)
	req, err := http.NewRequest("DELETE", "/users", payload)
	req.Header.Set(fiber.HeaderContentType, contentType)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestGetAllUsersValidateFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("GetAllUsers", mock.AnythingOfType(mockUserTypeSlice), "all", "asc", "id", "0", "10", "inbox").Return(int64(0), errors.New("error"))
	app := fiber.New()
	app.Get("/users", handler.GetAllUsers)
	req, err := http.NewRequest("GET", "/users?sort=hello", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestGetAllUsersFail(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("GetAllUsers", mock.AnythingOfType(mockUserTypeSlice), "all", "asc", "id", "0", "10", "inbox").Return(int64(0), errors.New("error"))
	app := fiber.New()
	app.Get("/users", handler.GetAllUsers)
	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestGetAllUsersSuccess(t *testing.T) {
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("GetAllUsers", mock.AnythingOfType(mockUserTypeSlice), "all", "asc", "id", "0", "10", "inbox").Return(int64(1), nil)
	app := fiber.New()
	app.Get("/users", handler.GetAllUsers)
	req, err := http.NewRequest("GET", "/users", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestRestoreUsersJSONFail(t *testing.T) {
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Put("/users", handler.RestoreUsers)
	req, err := http.NewRequest("PUT", "/users", nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRestoreUsersValidateFail(t *testing.T) {
	var uuid paramUUID
	data, _ := json.Marshal(&uuid)
	payload := bytes.NewReader(data)
	_, handler := UserMockSetup(t)
	app := fiber.New()
	app.Put("/users", handler.RestoreUsers)
	req, err := http.NewRequest("PUT", "/users", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestRestoreUsersRestoreFail(t *testing.T) {
	var uuid paramUUID
	uuid.UsersID = []string{
		UUID,
	}
	data, _ := json.Marshal(&uuid)
	payload := bytes.NewReader(data)
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("RestoreUsers", uuid.UsersID).Return(int64(0), errors.New("error TestRestoreUsersRestoreFail"))
	app := fiber.New()
	app.Put("/users", handler.RestoreUsers)
	req, err := http.NewRequest("PUT", "/users", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestRestoreUsersRestoreSuccess(t *testing.T) {
	var uuid paramUUID
	uuid.UsersID = []string{
		UUID,
	}
	data, _ := json.Marshal(&uuid)
	payload := bytes.NewReader(data)
	mockUseCase, handler := UserMockSetup(t)
	mockUseCase.On("RestoreUsers", uuid.UsersID).Return(int64(1), nil)
	app := fiber.New()
	app.Put("/users", handler.RestoreUsers)
	req, err := http.NewRequest("PUT", "/users", payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}

package deliveries

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/gofiber/fiber/v2"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func ForgotPasswordMockSetup(t *testing.T) (mockUseCase *mocks.ForgorPasswordUseCase, handler *forgotPasswordHandler) {
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
	mockUseCase = new(mocks.ForgorPasswordUseCase)
	handler = NewForgotPasswordHandler(mockUseCase, &mr)
	return
}
func TestResetPasswordJSONNotSend(t *testing.T) {
	_, handler := ForgotPasswordMockSetup(t)
	app := fiber.New()
	app.Put("/reset_password/:id", handler.ResetPassword)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/reset_password/"+userID, nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestResetPasswordValidateFail(t *testing.T) {
	_, handler := ForgotPasswordMockSetup(t)

	var params paramEmail
	err := faker.FakeData(&params)
	assert.NoError(t, err)
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	app := fiber.New()
	app.Put("/reset_password/:id", handler.ResetPassword)
	var userID string
	err = faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestResetPasswordCreateForgotPasswordHashFail(t *testing.T) {
	mockUseCase, handler := ForgotPasswordMockSetup(t)

	mockUseCase.On("CreateForgotPasswordHash", mock.Anything).Return("", errors.New("error CreateForgotPasswordHash"))
	params := paramEmail{
		Email: "aaa@aaa.com",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)

	app := fiber.New()
	app.Put("/reset_password/:id", handler.ResetPassword)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestResetPasswordMailHTMLError(t *testing.T) {
	databases.MockMailServer = mailgun.NewMockServer()
	defer databases.MockMailServer.Stop()

	mockUseCase, handler := ForgotPasswordMockSetup(t)

	mockUseCase.On("CreateForgotPasswordHash", mock.Anything).Return("", nil)
	mockUseCase.On("MailHTML", mock.Anything, mock.AnythingOfType("templateResetPassword")).Return("", errors.New("error MailHTML"))

	params := paramEmail{
		Email: "aaa@aaa.com",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/reset_password/:id", handler.ResetPassword)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 406, resp.StatusCode, "completed")
}
func TestResetPasswordSendMailError(t *testing.T) {
	databases.MockMailServer = mailgun.NewMockServer()
	defer databases.MockMailServer.Stop()

	mockUseCase, handler := ForgotPasswordMockSetup(t)

	mockUseCase.On("CreateForgotPasswordHash", mock.Anything).Return("", nil)
	mockUseCase.On("MailHTML", mock.Anything, mock.AnythingOfType("templateResetPassword")).Return("data", nil)
	mockUseCase.On("SendMail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error SendMail"))
	params := paramEmail{
		Email: "aaa@aaa.com",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/reset_password/:id", handler.ResetPassword)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 401, resp.StatusCode, "completed")
}
func TestResetPassword(t *testing.T) {
	databases.MockMailServer = mailgun.NewMockServer()
	defer databases.MockMailServer.Stop()

	mockUseCase, handler := ForgotPasswordMockSetup(t)

	mockUseCase.On("CreateForgotPasswordHash", mock.Anything).Return("", nil)
	mockUseCase.On("MailHTML", mock.Anything, mock.AnythingOfType("templateResetPassword")).Return("data", nil)
	mockUseCase.On("SendMail", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)
	params := paramEmail{
		Email: "aaa@aaa.com",
	}
	data, _ := json.Marshal(&params)
	payload := bytes.NewReader(data)
	app := fiber.New()
	app.Put("/reset_password/:id", handler.ResetPassword)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("PUT", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestResetPasswordFormCheckForgotPasswordHashIsExpireFail(t *testing.T) {
	mockUseCase, handler := ForgotPasswordMockSetup(t)

	mockUseCase.On("CheckForgotPasswordHashIsExpire", mock.Anything).Return(false)

	app := fiber.New()
	app.Get("/reset_password/:id", handler.ResetPasswordForm)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("GET", "/reset_password/"+userID, nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestResetPasswordForm(t *testing.T) {
	mockUseCase, handler := ForgotPasswordMockSetup(t)

	mockUseCase.On("CheckForgotPasswordHashIsExpire", mock.Anything).Return(true)

	app := fiber.New()
	app.Get("/reset_password/:id", handler.ResetPasswordForm)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("GET", "/reset_password/"+userID, nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestResetPasswordFormExecParamsFail(t *testing.T) {
	_, handler := ForgotPasswordMockSetup(t)

	app := fiber.New()
	app.Post("/reset_password/:id", handler.ResetPasswordFormExec)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/reset_password/"+userID, nil)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 400, resp.StatusCode, "completed")
}
func TestResetPasswordFormExecValidateFail(t *testing.T) {
	_, handler := ForgotPasswordMockSetup(t)
	data := url.Values{}
	data.Set("password", "pass")
	data.Set("re_password", "pass")
	payload := strings.NewReader(data.Encode())

	app := fiber.New()
	app.Post("/reset_password/:id", handler.ResetPasswordFormExec)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	req.Header.Set(fiber.HeaderContentLength, strconv.Itoa(len(data.Encode())))
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}
func TestResetPasswordFormExecResetPasswordFail(t *testing.T) {
	mockUseCase, handler := ForgotPasswordMockSetup(t)
	mockUseCase.On("ResetPassword", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("error ResetPassword"))

	data := url.Values{}
	data.Set("password", "password")
	data.Set("re_password", "password")
	payload := strings.NewReader(data.Encode())

	app := fiber.New()
	app.Post("/reset_password/:id", handler.ResetPasswordFormExec)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	req.Header.Set(fiber.HeaderContentLength, strconv.Itoa(len(data.Encode())))
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}

func TestResetPasswordFormExec(t *testing.T) {
	mockUseCase, handler := ForgotPasswordMockSetup(t)
	mockUseCase.On("ResetPassword", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	data := url.Values{}
	data.Set("password", "password")
	data.Set("re_password", "password")
	payload := strings.NewReader(data.Encode())

	app := fiber.New()
	app.Post("/reset_password/:id", handler.ResetPasswordFormExec)
	var userID string
	err := faker.FakeData(&userID)
	assert.NoError(t, err)
	req, err := http.NewRequest("POST", "/reset_password/"+userID, payload)
	assert.NoError(t, err)
	req.Header.Set(fiber.HeaderContentType, contentType)
	req.Header.Set(fiber.HeaderContentType, fiber.MIMEApplicationForm)
	req.Header.Set(fiber.HeaderContentLength, strconv.Itoa(len(data.Encode())))
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode, "completed")
}

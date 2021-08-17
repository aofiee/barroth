package usecases

import (
	"errors"
	"strings"
	"testing"

	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/mailgun/mailgun-go/v4"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateForgotPasswordHash(t *testing.T) {
	repo := new(mocks.ForgorPasswordRepository)
	var email string
	err := faker.FakeData(&email)
	assert.NoError(t, err)

	repo.On("CreateForgotPasswordHash", email, mock.AnythingOfType("string"), mock.AnythingOfType("time.Duration")).Return(nil).Once()

	u := NewForgotPasswordUseCase(repo)
	hash, err := u.CreateForgotPasswordHash(email)
	assert.NoError(t, err)
	assert.NotEqual(t, nil, hash)

	repo.On("GetHash", hash).Return(hash, nil).Once()
	ok := u.CheckForgotPasswordHashIsExpire(hash)
	assert.Equal(t, true, ok)
}

func TestResetPassword(t *testing.T) {
	repo := new(mocks.ForgorPasswordRepository)
	hash := ksuid.New().String()

	var password string
	err := faker.FakeData(&password)
	assert.NoError(t, err)
	var rePassword string
	err = faker.FakeData(&rePassword)
	assert.NoError(t, err)

	u := NewForgotPasswordUseCase(repo)
	err = u.ResetPassword(hash, password, rePassword)
	assert.Error(t, err)

	var email string
	err = faker.FakeData(&email)
	assert.NoError(t, err)

	rePassword = password
	repo.On("GetHash", hash).Return(string(""), errors.New("error GetHash")).Once()
	err = u.ResetPassword(hash, password, rePassword)
	assert.Error(t, err)

	var user models.Users
	err = faker.FakeData(&user)
	assert.NoError(t, err)

	repo.On("GetHash", hash).Return(email, nil).Once()
	repo.On("HashPassword", mock.AnythingOfType(userModelType)).Return(errors.New("error HashPassword")).Once()
	err = u.ResetPassword(hash, password, rePassword)
	assert.Error(t, err)

	repo.On("GetHash", hash).Return(email, nil).Once()
	repo.On("HashPassword", mock.AnythingOfType(userModelType)).Return(nil).Once()
	repo.On("ResetPassword", mock.AnythingOfType(userModelType)).Return(errors.New("error ResetPassword")).Once()
	err = u.ResetPassword(hash, password, rePassword)
	assert.Error(t, err)

	repo.On("GetHash", hash).Return(email, nil).Once()
	repo.On("HashPassword", mock.AnythingOfType(userModelType)).Return(nil).Once()
	repo.On("ResetPassword", mock.AnythingOfType(userModelType)).Return(nil).Once()
	repo.On("DeleteHash", mock.Anything).Return(nil).Once()
	err = u.ResetPassword(hash, password, rePassword)
	assert.NoError(t, err)
}
func TestMailHTML(t *testing.T) {
	repo := new(mocks.ForgorPasswordRepository)
	u := NewForgotPasswordUseCase(repo)
	type tmpstruct struct {
		AppName  string
		SiteURL  string
		LinkHash string
		AppPort  string
	}
	var tmp tmpstruct
	err := faker.FakeData(&tmp)
	assert.NoError(t, err)
	html, err := u.MailHTML("../views/change_password/mail_change_password.html", tmp)
	assert.NoError(t, err)
	assert.Equal(t, true, strings.Contains(html, tmp.AppName))
}
func TestSendMail(t *testing.T) {
	databases.MockMailServer = mailgun.NewMockServer()
	defer databases.MockMailServer.Stop()

	repo := new(mocks.ForgorPasswordRepository)
	u := NewForgotPasswordUseCase(repo)
	type datastruct struct {
		AppName  string
		SiteURL  string
		LinkHash string
		AppPort  string
	}
	var data datastruct
	err := faker.FakeData(&data)
	assert.NoError(t, err)
	html, err := u.MailHTML("../views/change_password/mail_change_password.html", data)
	assert.NoError(t, err)
	assert.Equal(t, true, strings.Contains(html, data.AppName))
	type tmpstruct struct {
		sender    string
		subject   string
		recipient string
		body      string
	}
	var tmp tmpstruct
	err = faker.FakeData(&tmp)
	assert.NoError(t, err)
	err = u.SendMail("domain", "key", "aaa@aaa.com", tmp.subject, "bbb@bbb.com", html)
	assert.NoError(t, err)
}

package usecases

import (
	"errors"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	authModelType = "*models.Users"
	authEmail     = "aofiee666@gmail.com"
	authPassword  = "password"
	authFullName  = "Arashi L."
	authTelephone = "0925905444"
)

func TestLoginSuccess(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)

	repo.On("Login", mock.AnythingOfType(authModelType), mock.Anything).Return(nil).Once()

	repo.On("CheckPasswordHash", mock.AnythingOfType(authModelType), mock.Anything).Return(true).Once()

	var user models.Users
	u := NewAuthenticationUseCase(repo)
	err := u.Login(&user, authEmail, authPassword)
	assert.NoError(t, err)
}
func TestLoginPasswordFail(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)

	repo.On("Login", mock.AnythingOfType(authModelType), mock.Anything).Return(nil).Once()

	repo.On("CheckPasswordHash", mock.AnythingOfType(authModelType), mock.Anything).Return(false).Once()

	var user models.Users
	u := NewAuthenticationUseCase(repo)
	err := u.Login(&user, authEmail, authPassword)
	assert.Error(t, err)
	assert.Equal(t, errors.New("username and password is incorrect"), err)
}
func TestLoginEmailNotFound(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)

	repo.On("Login", mock.AnythingOfType(authModelType), mock.Anything).Return(errors.New("record not found")).Once()

	var user models.Users
	u := NewAuthenticationUseCase(repo)
	err := u.Login(&user, authEmail, authPassword)
	assert.Error(t, err)
	assert.Equal(t, errors.New("record not found"), err)
}

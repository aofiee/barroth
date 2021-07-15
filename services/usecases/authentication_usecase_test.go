package usecases

import (
	"errors"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	authModelType              = "*models.Users"
	authTokenRoleNameModelType = "*models.TokenRoleName"
	authEmail                  = "aofiee666@gmail.com"
	authPassword               = "password"
	authFullName               = "Arashi L."
	authTelephone              = "0925905444"
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
func TestCreateTokenSuccess(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)
	repo.On("GetRoleNameByUserID", mock.AnythingOfType(authTokenRoleNameModelType), mock.AnythingOfType("uint")).Return(nil).Once()
	user := models.Users{
		UUID:      utils.UUIDv4(),
		Name:      authFullName,
		Email:     authEmail,
		Password:  authPassword,
		Telephone: authTelephone,
		Status:    0,
	}
	u := NewAuthenticationUseCase(repo)
	tokenDetail, err := u.CreateToken(&user)
	assert.NoError(t, err)
	assert.Equal(t, authEmail, tokenDetail.Context.Email)

	err = u.GenerateAccessTokenBy(&user, &tokenDetail)
	assert.NoError(t, err)
	assert.NotEqual(t, "", tokenDetail.Token.AccessToken)

	err = u.GenerateRefreshTokenBy(&user, &tokenDetail)
	assert.NoError(t, err)
	assert.NotEqual(t, "", tokenDetail.Token.RefreshToken)
}
func TestCreateTokenFail(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)
	repo.On("GetRoleNameByUserID", mock.AnythingOfType(authTokenRoleNameModelType), mock.AnythingOfType("uint")).Return(errors.New("get role name error")).Once()
	user := models.Users{
		UUID:      utils.UUIDv4(),
		Name:      authFullName,
		Email:     authEmail,
		Password:  authPassword,
		Telephone: authTelephone,
		Status:    0,
	}
	u := NewAuthenticationUseCase(repo)
	tokenDetail, err := u.CreateToken(&user)
	assert.Error(t, err)
	assert.Equal(t, "", tokenDetail.Context.Email)
}
func TestSaveTokenSuccess(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)
	repo.On("SaveToken", mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration")).Return(nil).Twice()

	var token models.TokenDetail
	u := NewAuthenticationUseCase(repo)
	err := u.SaveToken(utils.UUIDv4(), &token)
	assert.NoError(t, err)
}
func TestSaveTokenFailOne(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)
	repo.On("SaveToken", mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration")).Return(errors.New("error save token"))

	var token models.TokenDetail
	u := NewAuthenticationUseCase(repo)
	err := u.SaveToken(utils.UUIDv4(), &token)
	assert.Error(t, err)
}
func TestSaveTokenFailTwo(t *testing.T) {
	repo := new(mocks.AuthenticationRepository)
	repo.On("SaveToken", mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration")).Return(nil).Once()

	repo.On("SaveToken", mock.Anything, mock.Anything, mock.AnythingOfType("time.Duration")).Return(errors.New("error save token")).Once()
	var token models.TokenDetail
	u := NewAuthenticationUseCase(repo)
	err := u.SaveToken(utils.UUIDv4(), &token)
	assert.Error(t, err)
}

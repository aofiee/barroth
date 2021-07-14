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
	userModelType = "*models.Users"
	userEmail     = "Test@test.com"
	userPassword  = "password"
	userFullName  = "Arashi L."
	userTelephone = "0925905444"
)

func getUser(email, password, name, telephone string) models.Users {
	return models.Users{
		Email:     userEmail,
		Password:  userPassword,
		Name:      userFullName,
		Telephone: userTelephone,
	}
}
func TestCreateUser(t *testing.T) {
	repo := new(mocks.UserRepository)
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	repo.On("HashPassword", mock.AnythingOfType(userModelType)).Return(nil).Once()

	repo.On("GetUserByEmail", mock.AnythingOfType(userModelType), mock.Anything).Return(errors.New("record not found")).Once()

	repo.On("CreateUser", mock.AnythingOfType(userModelType)).Return(nil).Once()
	u := NewUserUseCase(repo)
	err := u.CreateUser(&user)
	assert.NoError(t, err)

	repo.On("HashPassword", mock.AnythingOfType(userModelType)).Return(nil).Once()

	repo.On("GetUserByEmail", mock.AnythingOfType(userModelType), mock.Anything).Return(nil).Once()

	repo.On("CreateUser", mock.AnythingOfType(userModelType)).Return(errors.New("email is duplicated")).Once()

	u = NewUserUseCase(repo)
	err = u.CreateUser(&user)
	assert.Error(t, err)

	repo.On("HashPassword", mock.AnythingOfType(userModelType)).Return(errors.New("hash error")).Once()

	u = NewUserUseCase(repo)
	err = u.CreateUser(&user)
	assert.Error(t, err)
}
func TestUpdateUserSuccess(t *testing.T) {
	repo := new(mocks.UserRepository)
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	repo.On("GetUser", mock.AnythingOfType(userModelType), mock.Anything).Return(nil).Once()
	repo.On("UpdateUser", mock.AnythingOfType(userModelType), mock.Anything).Return(nil).Once()
	u := NewUserUseCase(repo)
	err := u.UpdateUser(&user, "xxx")
	assert.NoError(t, err)
}

func TestUpdateUserFail(t *testing.T) {
	repo := new(mocks.UserRepository)
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	repo.On("GetUser", mock.AnythingOfType(userModelType), mock.Anything).Return(errors.New("error TestUpdateUserFail")).Once()
	u := NewUserUseCase(repo)
	err := u.UpdateUser(&user, "xxx")
	assert.Error(t, err)
}
func TestGetUser(t *testing.T) {
	repo := new(mocks.UserRepository)
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	repo.On("GetUser", mock.AnythingOfType(userModelType), mock.Anything).Return(nil).Once()
	u := NewUserUseCase(repo)
	err := u.GetUser(&user, "xxx")
	assert.NoError(t, err)
}
func TestGetAllUsers(t *testing.T) {
	var users []models.Users
	repo := new(mocks.UserRepository)
	repo.On("GetAllUsers", mock.AnythingOfType("*[]models.Users"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	u := NewUserUseCase(repo)
	err := u.GetAllUsers(&users, "all", "asc", "id", "1", "10", "inbox")
	assert.NoError(t, err)
}
func TestDeleteUsers(t *testing.T) {
	repo := new(mocks.UserRepository)
	repo.On("DeleteUsers", mock.Anything, mock.AnythingOfType("[]int")).Return(int64(3), nil).Once()
	u := NewUserUseCase(repo)
	rs, err := u.DeleteUsers("inbox", []int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, int64(3), rs)
}

func TestRestoreUsers(t *testing.T) {
	repo := new(mocks.UserRepository)
	repo.On("RestoreUsers", mock.AnythingOfType("[]int")).Return(int64(3), nil).Once()
	u := NewUserUseCase(repo)
	rs, err := u.RestoreUsers([]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, int64(3), rs)
}

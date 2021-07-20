package usecases

import (
	"errors"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	systemModelType = "*models.System"
	appName         = "Test"
	siteUrl         = "http://"
	isInstall       = 0
)

func TestCreateSystem(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   appName,
		IsInstall: isInstall,
		SiteURL:   siteUrl,
	}
	repo.On("CreateSystem", mock.AnythingOfType(systemModelType)).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.CreateSystem(&m)
	assert.NoError(t, err)
}
func TestUpdateSystemSuccess(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   appName,
		IsInstall: isInstall,
		SiteURL:   siteUrl,
	}
	repo.On("GetSystem", mock.AnythingOfType(systemModelType), mock.Anything).Return(nil).Once()
	repo.On("UpdateSystem", mock.AnythingOfType(systemModelType), mock.Anything).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.UpdateSystem(&m, "xx")
	assert.NoError(t, err)
}
func TestUpdateSystemFail(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   appName,
		IsInstall: isInstall,
		SiteURL:   siteUrl,
	}
	repo.On("GetSystem", mock.AnythingOfType(systemModelType), mock.Anything).Return(errors.New("error")).Once()
	u := NewSystemUseCase(repo)
	err := u.UpdateSystem(&m, "xx")
	assert.Error(t, err)
}
func TestGetSystem(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   appName,
		IsInstall: isInstall,
		SiteURL:   siteUrl,
	}
	repo.On("GetSystem", mock.AnythingOfType(systemModelType), mock.Anything).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.GetSystem(&m, "/test")
	assert.NoError(t, err)
}
func TestGetFirstSystemInstallation(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   appName,
		IsInstall: isInstall,
		SiteURL:   siteUrl,
	}
	repo.On("GetFirstSystemInstallation", mock.AnythingOfType(systemModelType)).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.GetFirstSystemInstallation(&m)
	assert.NoError(t, err)
}
func TestSystemCreateUser(t *testing.T) {
	repo := new(mocks.SystemRepository)
	repo.On("CreateUser", mock.AnythingOfType(userModelType)).Return(nil)
	u := NewSystemUseCase(repo)
	var user models.Users
	faker.FakeData(&user)
	err := u.CreateUser(&user)
	assert.NoError(t, err)
}
func TestSystemCreateRole(t *testing.T) {
	repo := new(mocks.SystemRepository)
	repo.On("CreateRole", mock.AnythingOfType(roleModelType)).Return(nil)
	u := NewSystemUseCase(repo)
	var role models.RoleItems
	faker.FakeData(&role)
	err := u.CreateRole(&role)
	assert.NoError(t, err)
}

package usecases

import (
	"errors"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateSystem(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   "test",
		IsInstall: 0,
		SiteURL:   "http://",
	}
	repo.On("CreateSystem", mock.AnythingOfType("*models.System")).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.CreateSystem(&m)
	assert.NoError(t, err)
}
func TestUpdateSystemSuccess(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   "test",
		IsInstall: 0,
		SiteURL:   "http://",
	}
	repo.On("GetSystem", mock.AnythingOfType("*models.System"), mock.Anything).Return(nil).Once()
	repo.On("UpdateSystem", mock.AnythingOfType("*models.System"), mock.Anything).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.UpdateSystem(&m, "xx")
	assert.NoError(t, err)
}
func TestUpdateSystemFail(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   "test",
		IsInstall: 0,
		SiteURL:   "http://",
	}
	repo.On("GetSystem", mock.AnythingOfType("*models.System"), mock.Anything).Return(errors.New("error")).Once()
	u := NewSystemUseCase(repo)
	err := u.UpdateSystem(&m, "xx")
	assert.Error(t, err)
}
func TestGetSystem(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   "test",
		IsInstall: 0,
		SiteURL:   "http://",
	}
	repo.On("GetSystem", mock.AnythingOfType("*models.System"), mock.Anything).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.GetSystem(&m, "/test")
	assert.NoError(t, err)
}
func TestGetFirstSystemInstallation(t *testing.T) {
	repo := new(mocks.SystemRepository)
	m := models.System{
		AppName:   "test",
		IsInstall: 0,
		SiteURL:   "http://",
	}
	repo.On("GetFirstSystemInstallation", mock.AnythingOfType("*models.System")).Return(nil).Once()
	u := NewSystemUseCase(repo)
	err := u.GetFirstSystemInstallation(&m)
	assert.NoError(t, err)
}

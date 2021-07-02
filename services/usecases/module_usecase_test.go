package usecases

import (
	"errors"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateModule(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        "test",
		Description: "test",
		ModuleSlug:  "/test",
	}
	repo.On("CreateModule", mock.AnythingOfType("*models.Modules")).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.CreateModule(&module)
	assert.NoError(t, err)
}

func TestUpdateModuleSuccess(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        "test",
		Description: "test",
		ModuleSlug:  "/test",
	}
	repo.On("GetModule", mock.AnythingOfType("*models.Modules"), mock.Anything).Return(nil).Once()
	repo.On("UpdateModule", mock.AnythingOfType("*models.Modules"), mock.Anything).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.UpdateModule(&module, "xx")
	assert.NoError(t, err)
}

func TestUpdateModuleFail(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        "test",
		Description: "test",
		ModuleSlug:  "/test",
	}
	repo.On("GetModule", mock.AnythingOfType("*models.Modules"), mock.Anything).Return(errors.New("error")).Once()
	u := NewModuleUseCase(repo)
	err := u.UpdateModule(&module, "xx")
	assert.Error(t, err)
}

func TestGetModule(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        "test",
		Description: "test",
		ModuleSlug:  "/test",
	}
	repo.On("GetModule", mock.AnythingOfType("*models.Modules"), mock.Anything).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.GetModule(&module, "/test")
	assert.NoError(t, err)
}

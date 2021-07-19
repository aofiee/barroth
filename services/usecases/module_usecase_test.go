package usecases

import (
	"errors"
	"testing"

	"github.com/aofiee/barroth/mocks"
	"github.com/aofiee/barroth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	moduleModelType = "*models.Modules"
	moduleName      = "Test"
	moduleDesc      = "Test"
)

func TestCreateModule(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	repo.On("CreateModule", mock.AnythingOfType(moduleModelType)).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.CreateModule(&module)
	assert.NoError(t, err)
}

func TestUpdateModuleSuccess(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	repo.On("GetModule", mock.AnythingOfType(moduleModelType), mock.Anything).Return(nil).Once()
	repo.On("UpdateModule", mock.AnythingOfType(moduleModelType), mock.Anything).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.UpdateModule(&module, "xx")
	assert.NoError(t, err)
}

func TestUpdateModuleFail(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	repo.On("GetModule", mock.AnythingOfType(moduleModelType), mock.Anything).Return(errors.New("error")).Once()
	u := NewModuleUseCase(repo)
	err := u.UpdateModule(&module, "xx")
	assert.Error(t, err)
}

func TestGetModule(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	repo.On("GetModule", mock.AnythingOfType(moduleModelType), mock.Anything).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.GetModule(&module, "/test")
	assert.NoError(t, err)
}
func TestGetModuleBySlug(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	repo.On("GetModuleBySlug", mock.AnythingOfType(moduleModelType), mock.Anything, mock.Anything).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.GetModuleBySlug(&module, fiber.MethodPost, "/test")
	assert.NoError(t, err)
}

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

func TestCreateModuleFail(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	repo.On("CreateModule", mock.AnythingOfType(moduleModelType)).Return(errors.New("error CreateModule")).Once()
	u := NewModuleUseCase(repo)
	err := u.CreateModule(&module)
	assert.Error(t, err)
}
func TestCreateModuleGetAllRolesFail(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	var roles []models.RoleItems
	roles = append(roles, models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	})
	repo.On("CreateModule", mock.AnythingOfType(moduleModelType)).Return(nil).Once()
	repo.On("GetAllRoles").Return(roles, errors.New("error GetAllRoles")).Once()
	u := NewModuleUseCase(repo)
	err := u.CreateModule(&module)
	assert.Error(t, err)
}
func TestCreateModuleSuccess(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	var roles []models.RoleItems
	roles = append(roles, models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	})
	repo.On("CreateModule", mock.AnythingOfType(moduleModelType)).Return(nil).Once()
	repo.On("GetAllRoles").Return(roles, nil).Once()
	repo.On("SetPermission", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.CreateModule(&module)
	assert.NoError(t, err)
}

func TestCreateModuleSuccessRoleAdministrator(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	module := models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
	}
	var roles []models.RoleItems
	roles = append(roles, models.RoleItems{
		Name:        "Administrator",
		Description: roleDesc,
	})
	repo.On("CreateModule", mock.AnythingOfType(moduleModelType)).Return(nil).Once()
	repo.On("GetAllRoles").Return(roles, nil).Once()
	repo.On("SetPermission", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil).Once()
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
func TestSetModulePermission(t *testing.T) {
	repo := new(mocks.ModuleRepository)
	repo.On("SetPermission", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.SetPermission(uint(1), uint(2), 0)
	assert.NoError(t, err)
}
func TestModuleGetAllRoles(t *testing.T) {
	var roles []models.RoleItems
	roles = append(roles, models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	})
	repo := new(mocks.ModuleRepository)
	repo.On("GetAllRoles").Return(roles, nil).Once()
	u := NewModuleUseCase(repo)

	a, err := u.GetAllRoles()
	assert.NoError(t, err)
	assert.Equal(t, a, roles)
}
func TestGetAllModules(t *testing.T) {
	var modules []models.Modules
	repo := new(mocks.ModuleRepository)
	modules = append(modules, models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
		Method:      fiber.MethodGet,
		ModuleSlug:  "/test",
	})
	repo.On("GetAllModules", mock.AnythingOfType("*[]models.Modules"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	u := NewModuleUseCase(repo)
	err := u.GetAllModules(&modules, "all", "asc", "id", "1", "10", "inbox")
	assert.NoError(t, err)
}

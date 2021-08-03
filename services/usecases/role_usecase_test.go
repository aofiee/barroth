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
	roleModelType = "*models.RoleItems"
	roleName      = "Test"
	roleDesc      = "Test"
)

func TestCreateRoleFail(t *testing.T) {
	repo := new(mocks.RoleRepository)
	role := models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	}
	repo.On("CreateRole", mock.AnythingOfType(roleModelType)).Return(errors.New("error CreateRole")).Once()
	u := NewRoleUseCase(repo)
	err := u.CreateRole(&role)
	assert.Error(t, err)
}
func TestCreateRoleGetAllModulesFail(t *testing.T) {
	repo := new(mocks.RoleRepository)
	role := models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	}
	repo.On("CreateRole", mock.AnythingOfType(roleModelType)).Return(nil).Once()
	var m []models.Modules
	repo.On("GetAllModules").Return(m, errors.New("error GetAllModules")).Once()
	u := NewRoleUseCase(repo)
	err := u.CreateRole(&role)
	assert.Error(t, err)
}
func TestCreateRoleSuccess(t *testing.T) {
	repo := new(mocks.RoleRepository)
	role := models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	}
	repo.On("CreateRole", mock.AnythingOfType(roleModelType)).Return(nil).Once()
	var m []models.Modules
	m = append(m, models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
		ModuleSlug:  "/test",
	})
	repo.On("GetAllModules").Return(m, nil).Once()
	repo.On("SetPermission", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil).Once()
	u := NewRoleUseCase(repo)
	err := u.CreateRole(&role)
	assert.NoError(t, err)
}
func TestUpdateRoleSuccess(t *testing.T) {
	repo := new(mocks.RoleRepository)
	role := models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	}
	repo.On("GetRole", mock.AnythingOfType(roleModelType), mock.Anything).Return(nil).Once()
	repo.On("UpdateRole", mock.AnythingOfType(roleModelType), mock.Anything).Return(nil).Once()
	u := NewRoleUseCase(repo)
	err := u.UpdateRole(&role, "xxx")
	assert.NoError(t, err)
}

func TestUpdateRoleFail(t *testing.T) {
	repo := new(mocks.RoleRepository)
	role := models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	}
	repo.On("GetRole", mock.AnythingOfType(roleModelType), mock.Anything).Return(errors.New("error")).Once()
	u := NewRoleUseCase(repo)
	err := u.UpdateRole(&role, "xxx")
	assert.Error(t, err)
}

func TestGetRole(t *testing.T) {
	repo := new(mocks.RoleRepository)
	role := models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	}
	repo.On("GetRole", mock.AnythingOfType(roleModelType), mock.Anything).Return(nil).Once()
	u := NewRoleUseCase(repo)
	err := u.GetRole(&role, "xxx")
	assert.NoError(t, err)
}
func TestGetAllRoles(t *testing.T) {
	repo := new(mocks.RoleRepository)
	var roles []models.RoleItems
	roles = append(roles, models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	})
	repo.On("GetAllRoles", mock.AnythingOfType("*[]models.RoleItems"), mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
	u := NewRoleUseCase(repo)
	err := u.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "inbox")
	assert.NoError(t, err)
}

func TestDeleteRoles(t *testing.T) {
	repo := new(mocks.RoleRepository)
	repo.On("DeleteRoles", mock.Anything, mock.AnythingOfType("[]int")).Return(int64(3), nil).Once()
	u := NewRoleUseCase(repo)
	rs, err := u.DeleteRoles("inbox", []int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, int64(3), rs)
}

func TestRestoreRoles(t *testing.T) {
	repo := new(mocks.RoleRepository)
	repo.On("RestoreRoles", mock.AnythingOfType("[]int")).Return(int64(3), nil).Once()
	u := NewRoleUseCase(repo)
	rs, err := u.RestoreRoles([]int{1, 2, 3})
	assert.NoError(t, err)
	assert.Equal(t, int64(3), rs)
}
func TestGetAllModule(t *testing.T) {
	var m []models.Modules
	m = append(m, models.Modules{
		Name:        moduleName,
		Description: moduleDesc,
		ModuleSlug:  "/test",
	})
	repo := new(mocks.RoleRepository)
	repo.On("GetAllModules").Return(m, nil).Once()
	u := NewRoleUseCase(repo)

	a, err := u.GetAllModules()
	assert.NoError(t, err)
	assert.Equal(t, a, m)
}
func TestSetRolePermission(t *testing.T) {
	repo := new(mocks.RoleRepository)
	repo.On("SetPermission", mock.AnythingOfType("uint"), mock.AnythingOfType("uint"), mock.AnythingOfType("int")).Return(nil).Once()
	u := NewRoleUseCase(repo)
	err := u.SetPermission(uint(1), uint(2), 0)
	assert.NoError(t, err)
}

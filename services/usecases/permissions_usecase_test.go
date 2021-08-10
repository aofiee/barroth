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
	permissionsModelType = "*models.Permissions"
)

func TestSetPermissionsNoValue(t *testing.T) {
	repo := new(mocks.PermissionsRepository)
	u := NewPermissionsUseCase(repo)
	var permissions []models.Permissions
	err := u.SetPermissions(&permissions)
	assert.NoError(t, err)
}
func TestSetPermissionsFail(t *testing.T) {
	repo := new(mocks.PermissionsRepository)
	repo.On("SetPermissions", mock.AnythingOfType(permissionsModelType)).Return(errors.New("error SetPermissions"))
	u := NewPermissionsUseCase(repo)
	var permissions []models.Permissions
	IsExec := 1
	permissions = append(permissions, models.Permissions{
		ModuleID:   1,
		RoleItemID: 1,
		IsExec:     &IsExec,
	})
	err := u.SetPermissions(&permissions)
	assert.Error(t, err)
}
func TestSetPermissionsSuccess(t *testing.T) {
	repo := new(mocks.PermissionsRepository)
	repo.On("SetPermissions", mock.AnythingOfType(permissionsModelType)).Return(nil)
	u := NewPermissionsUseCase(repo)
	var permissions []models.Permissions
	IsExec := 1
	permissions = append(permissions, models.Permissions{
		ModuleID:   1,
		RoleItemID: 1,
		IsExec:     &IsExec,
	})
	err := u.SetPermissions(&permissions)
	assert.NoError(t, err)
}

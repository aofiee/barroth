package repositories

import (
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
)

const (
	permissionsRepositoryType = "*repositories.PermissionsRepository"
)

func TestSetPermissionModule(t *testing.T) {
	SetupMock(t)
	repo := NewPermissionsRepository(databases.DB)
	assert.Equal(t, permissionsRepositoryType, reflect.TypeOf(repo).String(), "TestSetPermissionModule")
	mockIsExec := 1
	permission := models.Permissions{
		ModuleID:   1,
		RoleItemID: 1,
		IsExec:     &mockIsExec,
	}

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `permissions`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.SetPermissions(&permission)
	assert.NoError(t, err)
}

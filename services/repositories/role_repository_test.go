package repositories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
)

const (
	roleRepositoryType = "*repositories.roleRepository"
	roleName           = "Test"
	roleDesc           = "Description"
)

func getRole(name, desc string) models.RoleItems {
	return models.RoleItems{
		Name:        name,
		Description: desc,
	}
}
func TestCreateRole(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), "TestCreateRole")
	role := getRole(roleName, roleDesc)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `role_items` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateRole(&role)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `role_items` ").WillReturnError(errors.New("error TestCreateRole"))
	mock.ExpectCommit()
	err = repo.CreateRole(&role)
	assert.Error(t, err)
}

func TestGetRole(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), "TestGetRole")
	role := getRole(roleName, roleDesc)
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description"}

	mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, roleName, roleDesc))
	err := repo.GetRole(&role, "1")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").WithArgs("1").
		WillReturnError(errors.New("error TestGetRole"))
	err = repo.GetRole(&role, "1")
	assert.Error(t, err)
}

func TestGetAllRoles(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), "TestGetAllRoles")
	var roles []models.RoleItems
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description"}

	err := repo.GetAllRoles(&roles, "all", "asc", "id", "xx", "10", "inbox")
	assert.Error(t, err)

	err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "xx", "inbox")
	assert.Error(t, err)
	t.Run("TEST_INBOX", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, roleName, roleDesc))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error TestGetAllRoles Inbox "))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, roleName, roleDesc))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error TestGetAllRoles Inbox with Keyword"))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)
	})
	t.Run("TEST_TRASH", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, roleName, roleDesc))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error TestGetAllRoles Trash"))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, roleName, roleDesc))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error TestGetAllRoles Trash with Keyword"))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)
	})
}

func TestUpdateRole(t *testing.T) {
	SetupMock(t)
	// now := time.Now()
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), TestUpdateRole)
	role := getRole(roleName, roleDesc)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `role_items`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateRole(&role, "1")
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `role_items`").WillReturnError(errors.New("error TestUpdateRole"))
	mock.ExpectCommit()
	err = repo.UpdateRole(&role, "1")
	assert.Error(t, err)
}

func TestDeleteRoleSuccess(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), "TestDeleteRoleSuccess")

	t.Run("TEST_INBOX", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `role_items`").WillReturnResult(sqlmock.NewResult(1, 3))
		mock.ExpectCommit()
		_, err := repo.DeleteRoles("inbox", []int{1, 2, 3})
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `role_items`").WillReturnError(errors.New("error TestDeleteRoleSuccess"))
		mock.ExpectCommit()
		_, err = repo.DeleteRoles("inbox", []int{1, 2, 3})
		assert.Error(t, err)
	})

}
func TestDeleteRoleFail(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), "TestDeleteRoleFail")
	t.Run("TEST_TRASH", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `role_items`").WithArgs(1, 2, 3).WillReturnResult(sqlmock.NewResult(1, 3))
		mock.ExpectCommit()
		_, err := repo.DeleteRoles("trash", []int{1, 2, 3})
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `role_items`").WillReturnError(errors.New("error TestDeleteRoleFail"))
		mock.ExpectCommit()
		_, err = repo.DeleteRoles("trash", []int{1, 2, 3})
		assert.Error(t, err)
	})
}
func TestRestoreRoleSuccess(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, roleRepositoryType, reflect.TypeOf(repo).String(), "TestRestoreRoleSuccess")

	t.Run("TEST_RESTORE_SUCCESS", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `role_items`").WillReturnResult(sqlmock.NewResult(1, 3))
		mock.ExpectCommit()
		_, err := repo.RestoreRoles([]int{1, 2, 3})
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `role_items`").WillReturnError(errors.New("error TestRestoreRoleSuccess"))
		mock.ExpectCommit()
		_, err = repo.RestoreRoles([]int{1, 2, 3})
		assert.Error(t, err)
	})
}

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

func TestCreateRole(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, "*repositories.roleRepository", reflect.TypeOf(repo).String(), "new repo")
	role := models.RoleItems{
		Name:        "Test",
		Description: "Test",
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `role_items` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateRole(&role)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `role_items` ").WillReturnError(errors.New("error"))
	mock.ExpectCommit()
	err = repo.CreateRole(&role)
	assert.Error(t, err)
}

func TestGetRole(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, "*repositories.roleRepository", reflect.TypeOf(repo).String(), "new repo")
	role := models.RoleItems{
		Name:        "Test",
		Description: "Test",
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description"}

	mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "Desc"))
	err := repo.GetRole(&role, "1")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").WithArgs("1").
		WillReturnError(errors.New("error"))
	err = repo.GetRole(&role, "1")
	assert.Error(t, err)
}

func TestGetAllRoles(t *testing.T) {
	SetupMock(t)
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, "*repositories.roleRepository", reflect.TypeOf(repo).String(), "new repo")
	var roles []models.RoleItems
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description"}

	err := repo.GetAllRoles(&roles, "all", "asc", "id", "xx", "10", "inbox")
	assert.Error(t, err)

	err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "xx", "inbox")
	assert.Error(t, err)
	t.Run("TEST_INBOX", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "description"))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error"))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "description"))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error"))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)
	})
	t.Run("TEST_TRASH", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "description"))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error"))
		err = repo.GetAllRoles(&roles, "all", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "description"))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
			WillReturnError(errors.New("error"))
		err = repo.GetAllRoles(&roles, "Admin", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)
	})
}

func TestUpdateRole(t *testing.T) {
	SetupMock(t)
	// now := time.Now()
	repo := NewRoleRepository(databases.DB)
	assert.Equal(t, "*repositories.roleRepository", reflect.TypeOf(repo).String(), "new repo")
	role := models.RoleItems{
		Name:        "Test",
		Description: "Test",
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `role_items`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateRole(&role, "1")
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `role_items`").WillReturnError(errors.New("error"))
	mock.ExpectCommit()
	err = repo.UpdateRole(&role, "1")
	assert.Error(t, err)
}

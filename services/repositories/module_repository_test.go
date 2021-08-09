package repositories

import (
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

const (
	modelRepositoryType = "*repositories.moduleRepository"
	modelName           = "Test"
	modelDesc           = "Description"
)

func TestCreateModule(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestCreateModule")
	module := models.Modules{
		Name:        modelName,
		Description: modelName,
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `modules` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateModule(&module)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `modules` ").WillReturnError(errors.New("error TestCreateModule"))
	mock.ExpectCommit()
	err = repo.CreateModule(&module)
	assert.Error(t, err)
}

func TestUpdateModule(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestUpdateModule")
	module := models.Modules{
		Name:        modelName,
		Description: modelName,
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `modules`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateModule(&module, 1)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `modules`").WillReturnError(errors.New("error TestUpdateModule"))
	mock.ExpectCommit()
	err = repo.UpdateModule(&module, 1)
	assert.Error(t, err)
}

func TestGetModule(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestGetModule")
	module := models.Modules{
		Name:        modelName,
		Description: modelName,
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "module_slug"}

	mock.ExpectQuery("^SELECT (.+) FROM `modules`*").WithArgs(1).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, modelName, "Desc", "/url"))
	err := repo.GetModule(&module, 1)
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `modules`*").WithArgs(1).
		WillReturnError(errors.New("error TestGetModule"))
	err = repo.GetModule(&module, 1)
	assert.Error(t, err)
}
func TestGetModuleBySlug(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestGetModule")
	module := models.Modules{
		Name:        modelName,
		Description: modelName,
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "module_slug"}

	mock.ExpectQuery("^SELECT (.+) FROM `modules`*").WithArgs(fiber.MethodGet, "/test").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, modelName, "Desc", "/url"))
	err := repo.GetModuleBySlug(&module, fiber.MethodGet, "/test")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `modules`*").WithArgs("1").
		WillReturnError(errors.New("error TestGetModule"))
	err = repo.GetModuleBySlug(&module, fiber.MethodGet, "/test")
	assert.Error(t, err)
}
func TestModuleGetAllRolesItems(t *testing.T) {
	SetupMock(t)
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description"}
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestRoleGetAllModule")

	var roles []models.RoleItems
	roles = append(roles, models.RoleItems{
		Name:        roleName,
		Description: roleDesc,
	})

	mock.ExpectQuery("^SELECT (.+) FROM `role_items`*").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(0, nil, nil, nil, roleName, roleDesc))
	m, err := repo.GetAllRoles()
	assert.NoError(t, err)
	assert.Equal(t, roles, m)
}
func TestSetModulePermissionFail(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestRoleGetAllModule")
	mock.ExpectQuery("^SELECT (.+) FROM `permissions`*").WithArgs(uint(1), uint(1)).WillReturnError(errors.New("error TestSetPermission"))
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `permissions` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.SetPermission(1, 1, 0)
	assert.NoError(t, err)
}
func TestSetModulePermissionSuccess(t *testing.T) {
	SetupMock(t)
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "module_id", "role_item_id", "is_exec"}
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestRoleGetAllModule")
	mock.ExpectQuery("^SELECT (.+) FROM `permissions`*").WithArgs(uint(1), uint(1)).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, 1, 1, 0))
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `permissions`").WillReturnResult(sqlmock.NewResult(1, 3))
	mock.ExpectCommit()
	err := repo.SetPermission(uint(1), uint(1), 0)
	assert.NoError(t, err)
}
func TestGetAllModules(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, modelRepositoryType, reflect.TypeOf(repo).String(), "TestGetAllModules")
	var modules []models.Modules

	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "method", "module_slug"}

	err := repo.GetAllModules(&modules, "all", "asc", "id", "xx", "10", "inbox")
	assert.Error(t, err)

	err = repo.GetAllModules(&modules, "all", "asc", "id", "1", "xx", "inbox")
	assert.Error(t, err)

	t.Run("TEST_INBOX", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, modelName, modelDesc, fiber.MethodGet, "/test"))
		err = repo.GetAllModules(&modules, "all", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnError(errors.New("error GetAllModules Inbox "))
		err = repo.GetAllModules(&modules, "all", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, modelName, modelDesc, fiber.MethodGet, "/test"))
		err = repo.GetAllModules(&modules, "Admin", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnError(errors.New("error GetAllModules Keyword "))
		err = repo.GetAllModules(&modules, "Admin", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)
	})

	t.Run("TEST_TRASH", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, modelName, modelDesc, fiber.MethodGet, "/test"))
		err = repo.GetAllModules(&modules, "all", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnError(errors.New("error GetAllModules Inbox "))
		err = repo.GetAllModules(&modules, "all", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, modelName, modelDesc, fiber.MethodGet, "/test"))
		err = repo.GetAllModules(&modules, "Admin", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `modules`*").
			WillReturnError(errors.New("error GetAllModules Keyword "))
		err = repo.GetAllModules(&modules, "Admin", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)
	})
}

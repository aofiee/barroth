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

func TestCreateModule(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, "*repositories.moduleRepository", reflect.TypeOf(repo).String(), "new repo")
	module := models.Modules{
		Name:        "Test",
		Description: "Test",
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `modules` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateModule(&module)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `modules` ").WillReturnError(errors.New("error"))
	mock.ExpectCommit()
	err = repo.CreateModule(&module)
	assert.Error(t, err)
}

func TestUpdateModule(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, "*repositories.moduleRepository", reflect.TypeOf(repo).String(), "new repo")
	module := models.Modules{
		Name:        "Test",
		Description: "Test",
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `modules`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateModule(&module, "1")
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `modules`").WillReturnError(errors.New("error"))
	mock.ExpectCommit()
	err = repo.UpdateModule(&module, "1")
	assert.Error(t, err)
}

func TestGetModule(t *testing.T) {
	SetupMock(t)
	repo := NewModuleRepository(databases.DB)
	assert.Equal(t, "*repositories.moduleRepository", reflect.TypeOf(repo).String(), "new repo")
	module := models.Modules{
		Name:        "Test",
		Description: "Test",
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "name", "description", "module_slug"}

	mock.ExpectQuery("^SELECT (.+) FROM `modules`*").WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "Desc", "/url"))
	err := repo.GetModule(&module, "1")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `modules`*").WithArgs("1").
		WillReturnError(errors.New("error"))
	err = repo.GetModule(&module, "1")
	assert.Error(t, err)
}

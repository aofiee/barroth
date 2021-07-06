package repositories

import (
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *sql.DB
	mock sqlmock.Sqlmock
)

func SetupMock(t *testing.T) {
	var err error
	barroth_config.ENV, err = barroth_config.LoadConfig("../")
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
	db, mock, err = sqlmock.New()
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
	dial := mysql.New(mysql.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "mysql",
		Conn:       db,
	})
	rows := mock.NewRows([]string{"VERSION()"}).
		AddRow("5.7.34")
	assert.Equal(t, "*sqlmock.Rows", reflect.TypeOf(rows).String(), "new row")
	mock.ExpectQuery("SELECT VERSION()").
		WillReturnRows(rows)
	databases.DB, err = gorm.Open(dial, &gorm.Config{})
	if err != nil {
		assert.NotEqual(t, nil, err, err.Error())
	}
}
func TestCreateSystem(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, "*repositories.systemRepository", reflect.TypeOf(repo).String(), "new repo")
	sys := models.System{
		AppName:   "Test",
		SiteURL:   "http://",
		IsInstall: 0,
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `system` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateSystem(&sys)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `system` ").WillReturnError(errors.New("error"))
	mock.ExpectCommit()
	err = repo.CreateSystem(&sys)
	assert.Error(t, err)
}
func TestGetSystem(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, "*repositories.systemRepository", reflect.TypeOf(repo).String(), "new repo")
	sys := models.System{
		AppName:   "Test",
		SiteURL:   "http://",
		IsInstall: 0,
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "app_name", "site_url", "is_install"}

	mock.ExpectQuery("^SELECT (.+) FROM `system`*").WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "http://", 0))
	err := repo.GetSystem(&sys, "1")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `system`*").WithArgs("1").
		WillReturnError(errors.New("error"))
	err = repo.GetSystem(&sys, "1")
	assert.Error(t, err)
}
func TestUpdateSystem(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, "*repositories.systemRepository", reflect.TypeOf(repo).String(), "new repo")
	sys := models.System{
		AppName:   "Test",
		SiteURL:   "http://",
		IsInstall: 0,
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `system`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateSystem(&sys, "1")
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `system`").WillReturnError(errors.New("error"))
	mock.ExpectCommit()
	err = repo.UpdateSystem(&sys, "1")
	assert.Error(t, err)
}
func TestGetFirstSystemInstallation(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, "*repositories.systemRepository", reflect.TypeOf(repo).String(), "new repo")
	sys := models.System{
		AppName:   "Test",
		SiteURL:   "http://",
		IsInstall: 0,
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "app_name", "site_url", "is_install"}

	mock.ExpectQuery("^SELECT (.+) FROM `system` WHERE `system`.`deleted_at` IS NULL ORDER BY `system`.`id` LIMIT 1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, "Test", "http://", 0))
	err := repo.GetFirstSystemInstallation(&sys)
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `system` WHERE `system`.`deleted_at` IS NULL ORDER BY `system`.`id` LIMIT 1").
		WillReturnError(errors.New("error"))
	err = repo.GetFirstSystemInstallation(&sys)
	assert.Error(t, err)
}

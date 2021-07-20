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

const (
	systemRepositoryType = "*repositories.systemRepository"
	appName              = "Test"
	appUrl               = "http://"
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
	assert.Equal(t, "*sqlmock.Rows", reflect.TypeOf(rows).String(), "SetupMock")
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
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestCreateSystem")
	sys := models.System{
		AppName:   appName,
		SiteURL:   appUrl,
		IsInstall: 0,
	}
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `system` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateSystem(&sys)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `system` ").WillReturnError(errors.New("error TestCreateSystem"))
	mock.ExpectCommit()
	err = repo.CreateSystem(&sys)
	assert.Error(t, err)
}
func TestGetSystem(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestGetSystem")
	sys := models.System{
		AppName:   appName,
		SiteURL:   appUrl,
		IsInstall: 0,
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "app_name", "site_url", "is_install"}

	mock.ExpectQuery("^SELECT (.+) FROM `system`*").WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, appName, appUrl, 0))
	err := repo.GetSystem(&sys, "1")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `system`*").WithArgs("1").
		WillReturnError(errors.New("error TestGetSystem"))
	err = repo.GetSystem(&sys, "1")
	assert.Error(t, err)
}
func TestUpdateSystem(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestUpdateSystem")
	sys := models.System{
		AppName:   appName,
		SiteURL:   appUrl,
		IsInstall: 0,
	}
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `system`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateSystem(&sys, "1")
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `system`").WillReturnError(errors.New("error TestUpdateSystem"))
	mock.ExpectCommit()
	err = repo.UpdateSystem(&sys, "1")
	assert.Error(t, err)
}
func TestGetFirstSystemInstallation(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestGetFirstSystemInstallation")
	sys := models.System{
		AppName:   appName,
		SiteURL:   appUrl,
		IsInstall: 0,
	}
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "app_name", "site_url", "is_install"}

	mock.ExpectQuery("^SELECT (.+) FROM `system` WHERE `system`.`deleted_at` IS NULL ORDER BY `system`.`id` LIMIT 1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, appName, appUrl, 0))
	err := repo.GetFirstSystemInstallation(&sys)
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `system` WHERE `system`.`deleted_at` IS NULL ORDER BY `system`.`id` LIMIT 1").
		WillReturnError(errors.New("error TestGetFirstSystemInstallation"))
	err = repo.GetFirstSystemInstallation(&sys)
	assert.Error(t, err)
}
func TestSystemCreateUser(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestGetFirstSystemInstallation")

	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.CreateUser(&user)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `users` ").WillReturnError(errors.New("error TestCreateUser"))
	mock.ExpectCommit()
	err = repo.CreateUser(&user)
	assert.Error(t, err)
}
func TestSystemCreateRole(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestGetFirstSystemInstallation")
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
func TestSystemCheckPasswordHash(t *testing.T) {
	SetupMock(t)
	repo := NewSystemRepository(databases.DB)
	assert.Equal(t, systemRepositoryType, reflect.TypeOf(repo).String(), "TestSystemCheckPasswordHash")
	user := models.Users{
		Password: authPassword,
	}
	err := repo.HashPassword(&user)
	assert.NoError(t, err)
}

package repositories

import (
	"database/sql"
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
func TestSystemRepo(t *testing.T) {
	SetupMock(t)
	t.Run("TEST_SYSTEM_REPO", func(t *testing.T) {
		repo := NewSystemRepository(databases.DB)
		assert.Equal(t, "*repositories.systemRepository", reflect.TypeOf(repo).String(), "new repo")
		s := models.System{
			AppName:   "Test",
			IsInstall: 0,
			SiteURL:   "http://localhost:8181",
		}
		err := repo.CreateSystem(&s)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		err = repo.UpdateSystem(&s, "1")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		err = repo.GetFirstSystemInstallation(&s)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		err = repo.GetSystem(&s, "1")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
	})
}

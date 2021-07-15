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
	authenticationRepositoryType = "*repositories.authenticationRepository"
	authEmail                    = "aofiee666@gmail.com"
	authPassword                 = "password"
	authFullName                 = "Arashi L."
	authTelephone                = "0925905444"
	authID                       = uint(3)
	authRoleName                 = "RoleName"
)

func TestLogin(t *testing.T) {
	SetupMock(t)
	repo := NewAuthenticationRepository(databases.DB)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestLogin")

	columns := []string{"id", "created_at", "updated_at", "deleted_at", "email", "password", "name", "telephone", "image", "uuid", "status"}

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(authEmail).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, authEmail, authPassword, authFullName, authTelephone, "image", "uuid", 0))

	var user models.Users
	err := repo.Login(&user, authEmail)
	assert.NoError(t, err)
}

func TestCheckPasswordHash(t *testing.T) {
	SetupMock(t)
	repo := NewAuthenticationRepository(databases.DB)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestCheckPasswordHash")

	repoHash := NewUserRepository(databases.DB)
	user := models.Users{
		Password: authPassword,
	}
	err := repoHash.HashPassword(&user)
	assert.NoError(t, err)
	ok := repo.CheckPasswordHash(&user, authPassword)
	assert.Equal(t, true, ok)
}

func TestGetRoleNameByUserIDSuccess(t *testing.T) {
	SetupMock(t)
	repo := NewAuthenticationRepository(databases.DB)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestGetRoleNameByUserID")

	column := []string{"role_items.name"}
	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(authID).
		WillReturnRows(sqlmock.NewRows(column).AddRow(authRoleName))
	var role models.TokenRoleName
	err := repo.GetRoleNameByUserID(&role, authID)
	assert.NoError(t, err)
}
func TestGetRoleNameByUserIDFail(t *testing.T) {
	SetupMock(t)
	repo := NewAuthenticationRepository(databases.DB)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestGetRoleNameByUserID")

	column := []string{"role_items.name"}
	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(nil).
		WillReturnRows(sqlmock.NewRows(column).AddRow(authRoleName))
	var role models.TokenRoleName
	err := repo.GetRoleNameByUserID(&role, authID)
	assert.Error(t, err)
}

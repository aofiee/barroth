package repositories

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/go-redis/redismock/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
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
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
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
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
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
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
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
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestGetRoleNameByUserID")

	column := []string{"role_items.name"}
	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(nil).
		WillReturnRows(sqlmock.NewRows(column).AddRow(authRoleName))
	var role models.TokenRoleName
	err := repo.GetRoleNameByUserID(&role, authID)
	assert.Error(t, err)
}
func TestSaveToken(t *testing.T) {
	SetupMock(t)
	rd, mock := redismock.NewClientMock()
	databases.QueueClient = rd
	uuid := utils.UUIDv4()
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestSaveToken")

	now := time.Now()
	expire := time.Unix(120, 0)
	mock.ExpectSet(uuid, uuid, expire.Sub(now)).SetVal("OK")
	err := repo.SaveToken(uuid, uuid, expire.Sub(now))
	assert.NoError(t, err)
}
func TestDeleteToken(t *testing.T) {
	SetupMock(t)
	rd, redisMock := redismock.NewClientMock()
	databases.QueueClient = rd

	uuid := utils.UUIDv4()
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestDeleteToken")

	redisMock.ExpectDel(uuid).SetVal(0)
	err := repo.DeleteToken(uuid)
	assert.NoError(t, err)
}
func TestAuthenticationGetUserSuccess(t *testing.T) {
	SetupMock(t)
	rd, _ := redismock.NewClientMock()
	databases.QueueClient = rd
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestAuthenticationGetUser")
	UUID := utils.UUIDv4()
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "email", "password", "name", "telephone", "image", "uuid", "status"}

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(UUID).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", UUID, 0))
	var user models.Users
	err := repo.GetUser(&user, UUID)
	assert.NoError(t, err)
}
func TestAuthenticationGetUserFail(t *testing.T) {
	SetupMock(t)
	rd, _ := redismock.NewClientMock()
	databases.QueueClient = rd
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestAuthenticationGetUser")
	UUID := utils.UUIDv4()

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(UUID).
		WillReturnError(errors.New("error"))
	var user models.Users
	err := repo.GetUser(&user, UUID)
	assert.Error(t, err)
}
func TestGetAccessUUIDFromRedis(t *testing.T) {
	SetupMock(t)
	rd, mock := redismock.NewClientMock()
	databases.QueueClient = rd
	uuid := utils.UUIDv4()
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestGetAccessUUIDFromRedis")
	uuidResult := utils.UUIDv4()
	mock.ExpectGet(uuid).SetVal(uuidResult)
	result, err := repo.GetAccessUUIDFromRedis(uuid)
	assert.NoError(t, err)
	assert.Equal(t, uuidResult, result)
}
func TestCheckRoutePermission(t *testing.T) {
	SetupMock(t)
	repo := NewAuthenticationRepository(databases.DB, databases.QueueClient)
	assert.Equal(t, authenticationRepositoryType, reflect.TypeOf(repo).String(), "TestCheckRoutePermission")
	column := []string{"permissions.is_exec"}
	mock.ExpectQuery("^SELECT permissions.is_exec FROM *").WithArgs("RoleName", fiber.MethodGet, "/test").
		WillReturnRows(sqlmock.NewRows(column).AddRow(1))
	ok := repo.CheckRoutePermission("RoleName", fiber.MethodGet, "/test")
	assert.Equal(t, true, ok)
}

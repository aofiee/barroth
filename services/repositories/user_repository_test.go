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
	userRepositoryType = "*repositories.userRepository"
	userEmail          = "Test@test.com"
	userPassword       = "password"
	userFullName       = "Arashi L."
	userTelephone      = "0925905444"
)

func getUser(email, password, name, telephone string) models.Users {
	return models.Users{
		Email:     email,
		Password:  password,
		Name:      name,
		Telephone: telephone,
	}
}
func TestCreateUser(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestCreateUser")
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
func TestGetUser(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestGetUser")
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "email", "password", "name", "telephone", "image", "uuid", "status"}

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs("1").
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", "uuid", 0))
	err := repo.GetUser(&user, "1")
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs("1").
		WillReturnError(errors.New("error TestGetUser"))
	err = repo.GetUser(&user, "1")
	assert.Error(t, err)
}
func TestGetUserByEmail(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestGetUserByEmail")
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "email", "password", "name", "telephone", "image", "uuid", "status"}

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(userEmail).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", "uuid", 0))
	err := repo.GetUserByEmail(&user, userEmail)
	assert.NoError(t, err)

	mock.ExpectQuery("^SELECT (.+) FROM `users`*").WithArgs(userEmail).
		WillReturnError(errors.New("error TestGetUser"))
	err = repo.GetUserByEmail(&user, userEmail)
	assert.Error(t, err)
}
func TestGetAllUser(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestGetAllUser")
	var users []models.Users
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "email", "password", "name", "telephone", "image", "uuid", "status"}

	rows, err := repo.GetAllUsers(&users, "all", "asc", "id", "xx", "10", "inbox")
	assert.Error(t, err)
	assert.Equal(t, int64(0), rows)
	rows, err = repo.GetAllUsers(&users, "all", "asc", "id", "1", "xx", "inbox")
	assert.Error(t, err)
	t.Run("TEST_INBOX", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", "uuid", 0))
		rows, err = repo.GetAllUsers(&users, "all", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnError(errors.New("error TestGetAllUser Inbox "))
		rows, err = repo.GetAllUsers(&users, "all", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", "uuid", 0))
		rows, err = repo.GetAllUsers(&users, "Admin", "asc", "id", "1", "10", "inbox")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnError(errors.New("error TestGetAllUser Inbox with Keyword"))
		rows, err = repo.GetAllUsers(&users, "Admin", "asc", "id", "1", "10", "inbox")
		assert.Error(t, err)
	})
	t.Run("TEST_TRASH", func(t *testing.T) {
		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", "uuid", 0))
		rows, err = repo.GetAllUsers(&users, "all", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnError(errors.New("error TestGetAllUser Trash"))
		rows, err = repo.GetAllUsers(&users, "all", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, userEmail, userPassword, userFullName, userTelephone, "image", "uuid", 0))
		rows, err = repo.GetAllUsers(&users, "Admin", "asc", "id", "1", "10", "trash")
		assert.NoError(t, err)

		mock.ExpectQuery("^SELECT (.+) FROM `users`*").
			WillReturnError(errors.New("error TestGetAllUser Trash with Keyword"))
		rows, err = repo.GetAllUsers(&users, "Admin", "asc", "id", "1", "10", "trash")
		assert.Error(t, err)
	})
}
func TestUpdateUser(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestUpdateUser")
	user := getUser(userEmail, userPassword, userFullName, userTelephone)
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.UpdateUser(&user, "1")
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").WillReturnError(errors.New("error TestUpdateUser"))
	mock.ExpectCommit()
	err = repo.UpdateUser(&user, "1")
	assert.Error(t, err)
}
func TestDeleteUserSuccess(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestDeleteUserSuccess")

	t.Run("TEST_INBOX", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users`").WillReturnResult(sqlmock.NewResult(1, 2))
		mock.ExpectCommit()
		_, err := repo.DeleteUsers("inbox", []string{"xxxx", "yyyyy"})
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users`").WillReturnError(errors.New("error TestDeleteUserSuccess"))
		mock.ExpectCommit()
		_, err = repo.DeleteUsers("inbox", []string{"xxxx", "yyyyy"})
		assert.Error(t, err)
	})

}
func TestDeleteUserFail(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestDeleteUserFail")
	t.Run("TEST_TRASH", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `users`").WithArgs("xxxx", "yyyyy").WillReturnResult(sqlmock.NewResult(1, 2))
		mock.ExpectCommit()
		_, err := repo.DeleteUsers("trash", []string{"xxxx", "yyyyy"})
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `users`").WillReturnError(errors.New("error TestDeleteUserFail"))
		mock.ExpectCommit()
		_, err = repo.DeleteUsers("trash", []string{"xxxx", "yyyyy"})
		assert.Error(t, err)
	})
}
func TestRestoreUserSuccess(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestRestoreUserSuccess")

	t.Run("TEST_RESTORE_SUCCESS", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users`").WillReturnResult(sqlmock.NewResult(1, 3))
		mock.ExpectCommit()
		_, err := repo.RestoreUsers([]string{"1", "2", "3"})
		assert.NoError(t, err)

		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `users`").WillReturnError(errors.New("error TestRestoreUserSuccess"))
		mock.ExpectCommit()
		_, err = repo.RestoreUsers([]string{"1", "2", "3"})
		assert.Error(t, err)
	})
}

func TestGetUserRole(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestGetUserRole")
	columns := []string{"id", "created_at", "updated_at", "deleted_at", "role_item_id", "user_id"}
	mock.ExpectQuery("^SELECT (.+) FROM `user_roles`*").WithArgs(uint(1)).
		WillReturnRows(sqlmock.NewRows(columns).AddRow(1, nil, nil, nil, uint(1), uint(1)))

	err := repo.GetUserRole(uint(1))
	assert.NoError(t, err)
}
func TestUpdateUserRole(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestUpdateUserRole")
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `user_roles`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	role := models.UserRoles{
		RoleItemID: 1,
		UserID:     1,
	}
	err := repo.UpdateUserRole(&role, uint(1))
	assert.NoError(t, err)
}
func TestCreateUserRole(t *testing.T) {
	SetupMock(t)
	repo := NewUserRepository(databases.DB)
	assert.Equal(t, userRepositoryType, reflect.TypeOf(repo).String(), "TestCreateUserRole")
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `user_roles` ").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	role := models.UserRoles{
		RoleItemID: 1,
		UserID:     1,
	}
	err := repo.CreateUserRole(&role)
	assert.NoError(t, err)
}

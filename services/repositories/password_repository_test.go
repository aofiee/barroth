package repositories

import (
	"reflect"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/bxcodec/faker"
	"github.com/go-redis/redismock/v8"
	"github.com/segmentio/ksuid"
	"github.com/stretchr/testify/assert"
)

const (
	forgotPasswordRepositoryType = "*repositories.forgotPasswordRepository"
)

func TestCreateForgotPasswordHash(t *testing.T) {
	SetupMock(t)
	rd, mock := redismock.NewClientMock()
	databases.ResetPasswordQueueClient = rd

	repo := NewForgotPasswordRepository(databases.DB, databases.ResetPasswordQueueClient)
	assert.Equal(t, forgotPasswordRepositoryType, reflect.TypeOf(repo).String(), "TestCreateForgotPasswordHash")

	now := time.Now()
	expire := time.Unix(120, 0)
	hash := ksuid.New()
	var email string
	err := faker.FakeData(&email)
	assert.NoError(t, err)
	mock.ExpectSet(hash.String(), email, expire.Sub(now)).SetVal("OK")
	err = repo.CreateForgotPasswordHash(email, hash.String(), expire.Sub(now))
	assert.NoError(t, err)

	mock.ExpectGet(hash.String()).SetVal(hash.String())
	rs, err := repo.GetHash(hash.String())
	assert.NoError(t, err)
	assert.Equal(t, hash.String(), rs)

	mock.ExpectDel(hash.String()).SetVal(0)
	err = repo.DeleteHash(hash.String())
	assert.NoError(t, err)

	var user models.Users
	err = faker.FakeData(&user)
	noneHash := user.Password
	assert.NoError(t, err)
	err = repo.HashPassword(&user)
	assert.NoError(t, err)
	assert.NotEqual(t, noneHash, user.Password)
}

func TestResetPassword(t *testing.T) {
	SetupMock(t)
	repo := NewForgotPasswordRepository(databases.DB, databases.ResetPasswordQueueClient)
	assert.Equal(t, forgotPasswordRepositoryType, reflect.TypeOf(repo).String(), "TestResetPassword")
	var user models.Users
	err := faker.FakeData(&user)
	assert.NoError(t, err)

	mock.ExpectBegin()
	mock.ExpectExec("UPDATE `users`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = repo.ResetPassword(&user)
	assert.NoError(t, err)
}

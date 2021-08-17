package repositories

import (
	"context"
	"time"

	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	forgotPasswordRepository struct {
		conn  *gorm.DB
		queue *redis.Client
	}
)

func NewForgotPasswordRepository(conn *gorm.DB, queue *redis.Client) domains.ForgorPasswordRepository {
	return &forgotPasswordRepository{
		conn:  conn,
		queue: queue,
	}
}

func (f *forgotPasswordRepository) CreateForgotPasswordHash(email, hash string, expire time.Duration) error {
	ctx := context.Background()
	err := f.queue.Set(ctx, hash, email, expire).Err()
	return err
}
func (f *forgotPasswordRepository) GetHash(hash string) (string, error) {
	var ctx = context.Background()
	email, err := f.queue.Get(ctx, hash).Result()
	return email, err
}
func (f *forgotPasswordRepository) DeleteHash(hash string) error {
	var ctx = context.Background()
	err := f.queue.Del(ctx, hash).Err()
	return err
}
func (f *forgotPasswordRepository) ResetPassword(m *models.Users) error {
	err := f.conn.Model(models.Users{}).Omit("id", "uuid", "email", "name", "telephone", "image", "uuid", "status").Where("email = ?", m.Email).Updates(m).Error
	return err
}
func (f *forgotPasswordRepository) HashPassword(user *models.Users) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	return err
}

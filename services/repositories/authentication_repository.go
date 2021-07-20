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
	authenticationRepository struct {
		conn  *gorm.DB
		queue *redis.Client
	}
)

func NewAuthenticationRepository(conn *gorm.DB, queue *redis.Client) domains.AuthenticationRepository {
	return &authenticationRepository{
		conn:  conn,
		queue: queue,
	}
}
func (a *authenticationRepository) Login(m *models.Users, email string) error {
	u := NewUserRepository(a.conn)
	err := u.GetUserByEmail(m, email)
	return err
}
func (a *authenticationRepository) CheckPasswordHash(m *models.Users, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}
func (a *authenticationRepository) GetRoleNameByUserID(m *models.TokenRoleName, id uint) error {
	rs := a.conn.Model(&models.Users{}).Select("role_items.name").Joins("inner join user_roles on users.id = user_roles.user_id").Joins("inner join role_items on role_items.id = user_roles.role_item_id").Where("users.id = ?", id).Find(&m)

	if rs.Error != nil {
		return rs.Error
	}
	return nil
}
func (a *authenticationRepository) SaveToken(uuid string, t string, expire time.Duration) error {
	ctx := context.Background()
	err := a.queue.Set(ctx, t, uuid, expire).Err()
	return err
}
func (a *authenticationRepository) DeleteToken(uuid string) error {
	var ctx = context.Background()
	err := a.queue.Del(ctx, uuid).Err()
	return err
}
func (a *authenticationRepository) GetUser(m *models.Users, uuid string) error {
	if err := a.conn.Where("uuid = ?", uuid).First(m).Error; err != nil {
		return err
	}
	return nil
}
func (a *authenticationRepository) GetAccessUUIDFromRedis(uuid string) (string, error) {
	ctx := context.Background()
	result, err := a.queue.Get(ctx, uuid).Result()
	return result, err
}

package repositories

import (
	"context"
	"log"
	"time"

	"github.com/aofiee/barroth/domains"
	"github.com/go-redis/redis/v8"
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
	log.Println("hash, email, expire", hash, email, expire)
	err := f.queue.Set(ctx, hash, email, expire).Err()
	return err
}

package domains

import (
	"time"

	"github.com/aofiee/barroth/models"
)

type (
	ForgorPasswordUseCase interface {
		CreateForgotPasswordHash(email string) (hash string, err error)
		CheckForgotPasswordHashIsExpire(hash string) (ok bool)
		ResetPassword(hash, password, rePassword string) (err error)
	}
	ForgorPasswordRepository interface {
		CreateForgotPasswordHash(email, hash string, expire time.Duration) (err error)
		GetHash(hash string) (email string, err error)
		DeleteHash(hash string) (err error)
		ResetPassword(m *models.Users) (err error)
		HashPassword(user *models.Users) (err error)
	}
)

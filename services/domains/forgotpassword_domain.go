package domains

import "time"

type (
	ForgorPasswordUseCase interface {
		CreateForgotPasswordHash(email string) (hash string, err error)
	}
	ForgorPasswordRepository interface {
		CreateForgotPasswordHash(email, hash string, expire time.Duration) (err error)
	}
)

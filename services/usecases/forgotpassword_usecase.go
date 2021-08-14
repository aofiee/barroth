package usecases

import (
	"time"

	"github.com/aofiee/barroth/domains"
	"github.com/segmentio/ksuid"
)

type (
	forgotPasswordUserCase struct {
		forgotPasswordRepo domains.ForgorPasswordRepository
	}
)

func NewForgotPasswordUseCase(repo domains.ForgorPasswordRepository) domains.ForgorPasswordUseCase {
	return &forgotPasswordUserCase{
		forgotPasswordRepo: repo,
	}
}

func (f *forgotPasswordUserCase) CreateForgotPasswordHash(email string) (string, error) {
	loc, _ := time.LoadLocation(timeLoc)
	now := time.Now().In(loc)
	expireIn := time.Now().In(loc).Add(time.Minute * 24).Unix()
	linkExpire := time.Unix(expireIn, 0).In(loc)
	hash := ksuid.New()
	err := f.forgotPasswordRepo.CreateForgotPasswordHash(email, hash.String(), linkExpire.Sub(now))
	return hash.String(), err
}

package usecases

import (
	"errors"
	"time"

	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
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
	expireIn := time.Now().In(loc).Add(time.Hour * 24).Unix()
	linkExpire := time.Unix(expireIn, 0).In(loc)
	hash := ksuid.New()
	err := f.forgotPasswordRepo.CreateForgotPasswordHash(email, hash.String(), linkExpire.Sub(now))
	return hash.String(), err
}
func (f *forgotPasswordUserCase) CheckForgotPasswordHashIsExpire(hash string) bool {
	_, err := f.forgotPasswordRepo.GetHash(hash)
	return err == nil
}
func (f *forgotPasswordUserCase) ResetPassword(hash, password, rePassword string) error {
	if password != rePassword {
		return errors.New(constants.ERR_PASSWORD_DOES_NOT_MATCH)
	}
	email, err := f.forgotPasswordRepo.GetHash(hash)
	if err != nil {
		return err
	}
	var user models.Users
	user.Email = email
	user.Password = password
	err = f.forgotPasswordRepo.HashPassword(&user)
	if err != nil {
		return err
	}
	err = f.forgotPasswordRepo.ResetPassword(&user)
	if err != nil {
		return err
	}
	err = f.forgotPasswordRepo.DeleteHash(hash)
	return err
}

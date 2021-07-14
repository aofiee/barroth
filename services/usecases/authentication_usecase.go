package usecases

import (
	"errors"

	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
)

type (
	authenticationUseCase struct {
		authenticationRepo domains.AuthenticationRepository
	}
)

func NewAuthenticationUseCase(repo domains.AuthenticationRepository) domains.AuthenticationUseCase {
	return &authenticationUseCase{
		authenticationRepo: repo,
	}
}
func (a *authenticationUseCase) Login(m *models.Users, email, password string) error {
	err := a.authenticationRepo.Login(m, email)
	if err != nil {
		return err
	}
	ok := a.authenticationRepo.CheckPasswordHash(m, password)
	if !ok {
		return errors.New(constants.ERR_USERNAME_PASSWORD_INCORRECT)
	}
	return nil
}

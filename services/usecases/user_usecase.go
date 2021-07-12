package usecases

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
)

type (
	userUseCase struct {
		userRepo domains.UserRepository
	}
)

func NewUserUseCase(repo domains.UserRepository) domains.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}
func (u *userUseCase) CreateUser(user *models.Users) error {
	err := u.userRepo.CreateUser(user)
	return err
}
func (u *userUseCase) UpdateUser(user *models.Users, uuid string) error {
	var chk models.Users
	err := u.userRepo.GetUser(&chk, uuid)
	if err != nil {
		return err
	}
	err = u.userRepo.UpdateUser(user, uuid)
	return err
}
func (u *userUseCase) GetUser(user *models.Users, uuid string) error {
	err := u.userRepo.GetUser(user, uuid)
	return err
}
func (u *userUseCase) GetAllUsers(user *[]models.Users, keyword, sorting, sortField, page, limit, focus string) error {
	err := u.userRepo.GetAllUsers(user, keyword, sorting, sortField, page, limit, focus)
	return err
}
func (u *userUseCase) DeleteUsers(focus string, uuid []int) (int64, error) {
	rs, err := u.userRepo.DeleteUsers(focus, uuid)
	return rs, err
}
func (u *userUseCase) RestoreUsers(uuid []int) (int64, error) {
	rs, err := u.userRepo.RestoreUsers(uuid)
	return rs, err
}

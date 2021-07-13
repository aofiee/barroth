package domains

import "github.com/aofiee/barroth/models"

type (
	UserUseCase interface {
		CreateUser(m *models.Users) (err error)
		GetUser(m *models.Users, id string) (err error)
		GetAllUsers(m *[]models.Users, keyword, sorting, sortField, page, limit, focus string) (err error)
		UpdateUser(m *models.Users, id string) (err error)
		DeleteUsers(focus string, id []int) (rs int64, err error)
		RestoreUsers(id []int) (rs int64, err error)
		HashPassword(user *models.Users) (err error)
		CheckPasswordHash(user *models.Users, password string) (ok bool)
	}
	UserRepository interface {
		CreateUser(m *models.Users) (err error)
		GetUser(m *models.Users, id string) (err error)
		GetUserByEmail(m *models.Users, email string) (err error)
		GetAllUsers(m *[]models.Users, keyword, sorting, sortField, page, limit, focus string) (err error)
		UpdateUser(m *models.Users, id string) (err error)
		DeleteUsers(focus string, id []int) (rs int64, err error)
		RestoreUsers(id []int) (rs int64, err error)
	}
)

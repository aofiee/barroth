package domains

import "github.com/aofiee/barroth/models"

type (
	UserUseCase interface {
		CreateUser(m *models.Users) (err error)
		GetUser(m *models.Users, id string) (err error)
		GetAllUser(m *[]models.Users, keyword, sorting, sortField, page, limit, focus string) (err error)
		UpdateUser(m *models.Users, id string) (err error)
		DeleteUsers(focus string, id []int) (rs int64, err error)
		RestoreUsers(id []int) (rs int64, err error)
	}
	UserRepository interface {
		CreateUser(m *models.Users) (err error)
		GetUser(m *models.Users, id string) (err error)
		GetAllUser(m *[]models.Users, keyword, sorting, sortField, page, limit, focus string) (err error)
		UpdateUser(m *models.Users, id string) (err error)
		DeleteUsers(focus string, id []int) (rs int64, err error)
		RestoreUsers(id []int) (rs int64, err error)
	}
)

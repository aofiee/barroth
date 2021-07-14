package domains

import "github.com/aofiee/barroth/models"

type (
	AuthenticationUseCase interface {
		Login(m *models.Users, email, password string) (err error)
	}
	AuthenticationRepository interface {
		Login(m *models.Users, email string) (err error)
		CheckPasswordHash(m *models.Users, password string) (ok bool)
	}
)

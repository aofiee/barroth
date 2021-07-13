package domains

import "github.com/aofiee/barroth/models"

type (
	AuthenticationUseCase interface {
		Login(m *models.Users, username, password string) (err error)
		Logout() (err error)
		RefreshToken(ExpireRefreshToken string) (token string, err error)
	}
	AuthenticationRepository interface {
		Login(m *models.Users, username, password string) (err error)
		Logout() (err error)
		RefreshToken(ExpireRefreshToken string) (token string, err error)
	}
)

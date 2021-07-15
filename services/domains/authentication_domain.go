package domains

import "github.com/aofiee/barroth/models"

type (
	AuthenticationUseCase interface {
		Login(m *models.Users, email, password string) (err error)
		CreateToken(m *models.Users) (token models.TokenDetail, err error)
		GenerateAccessTokenBy(u *models.Users, t *models.TokenDetail) (err error)
		GenerateRefreshTokenBy(u *models.Users, t *models.TokenDetail) (err error)
	}
	AuthenticationRepository interface {
		Login(m *models.Users, email string) (err error)
		CheckPasswordHash(m *models.Users, password string) (ok bool)
		GetRoleNameByUserID(m *models.TokenRoleName, id uint) (err error)
	}
)

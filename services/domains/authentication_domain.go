package domains

import (
	"time"

	"github.com/aofiee/barroth/models"
)

type (
	AuthenticationUseCase interface {
		Login(m *models.Users, email, password string) (err error)
		CreateToken(m *models.Users) (token models.TokenDetail, err error)
		GenerateAccessTokenBy(u *models.Users, t *models.TokenDetail) (err error)
		GenerateRefreshTokenBy(u *models.Users, t *models.TokenDetail) (err error)
		SaveToken(uuid string, t *models.TokenDetail) (err error)
		DeleteToken(uuid string) (err error)
		GetUser(m *models.Users, uuid string) (err error)
	}
	AuthenticationRepository interface {
		Login(m *models.Users, email string) (err error)
		CheckPasswordHash(m *models.Users, password string) (ok bool)
		GetRoleNameByUserID(m *models.TokenRoleName, id uint) (err error)
		SaveToken(uuid string, tokenUUID string, expire time.Duration) (err error)
		DeleteToken(uuid string) (err error)
		GetUser(m *models.Users, uuid string) (err error)
	}
)

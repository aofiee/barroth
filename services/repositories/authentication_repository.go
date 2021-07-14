package repositories

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	authenticationRepository struct {
		conn *gorm.DB
	}
)

func NewAuthenticationRepository(conn *gorm.DB) domains.AuthenticationRepository {
	return &authenticationRepository{conn}
}
func (a *authenticationRepository) Login(m *models.Users, email string) error {
	u := NewUserRepository(a.conn)
	err := u.GetUserByEmail(m, email)
	return err
}
func (a *authenticationRepository) CheckPasswordHash(m *models.Users, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password))
	return err == nil
}

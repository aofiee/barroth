package repositories

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	systemRepository struct {
		conn *gorm.DB
	}
)

func NewSystemRepository(conn *gorm.DB) domains.SystemRepository {
	return &systemRepository{conn}
}
func (s *systemRepository) GetSystem(system *models.System, id string) error {
	if err := s.conn.Where("id = ?", id).First(system).Error; err != nil {
		return err
	}
	return nil
}
func (s *systemRepository) CreateSystem(system *models.System) error {
	if err := s.conn.Create(system).Error; err != nil {
		return err
	}
	return nil
}
func (s *systemRepository) UpdateSystem(system *models.System, id string) error {
	if err := s.conn.Model(system).Omit("id").Where("id = ?").Updates(system).Error; err != nil {
		return err
	}
	return nil
}
func (s *systemRepository) GetFirstSystemInstallation(system *models.System) error {
	if err := s.conn.First(system).Error; err != nil {
		return err
	}
	return nil
}
func (s *systemRepository) CreateUser(m *models.Users) error {
	if err := s.conn.Create(m).Error; err != nil {
		return err
	}
	return nil
}
func (s *systemRepository) CreateRole(m *models.RoleItems) error {
	if err := s.conn.Create(m).Error; err != nil {
		return err
	}
	return nil
}
func (s *systemRepository) HashPassword(m *models.Users) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(m.Password), 14)
	m.Password = string(bytes)
	return err
}

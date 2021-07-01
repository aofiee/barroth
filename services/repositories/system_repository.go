package repositories

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
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
	if err := s.conn.Save(system).Error; err != nil {
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

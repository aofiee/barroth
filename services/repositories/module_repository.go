package repositories

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"gorm.io/gorm"
)

type (
	moduleRepository struct {
		conn *gorm.DB
	}
)

func NewModuleRepository(conn *gorm.DB) domains.ModuleRepository {
	return &moduleRepository{conn}
}
func (m *moduleRepository) GetModuleBySlug(module *models.Modules, method, slug string) error {
	if err := m.conn.Where("method = ? AND module_slug = ?", method, slug).First(module).Error; err != nil {
		return err
	}
	return nil
}
func (m *moduleRepository) GetModule(module *models.Modules, id string) error {
	if err := m.conn.Where("ID = ?", id).First(module).Error; err != nil {
		return err
	}
	return nil
}
func (m *moduleRepository) CreateModule(module *models.Modules) error {
	if err := m.conn.Create(module).Error; err != nil {
		return err
	}
	return nil
}
func (m *moduleRepository) UpdateModule(module *models.Modules, id string) error {
	if err := m.conn.Model(module).Omit("id").Where("id = ?").Updates(module).Error; err != nil {
		return err
	}
	return nil
}

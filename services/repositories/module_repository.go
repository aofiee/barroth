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
	err := m.conn.Where("ID = ?", id).First(module).Error
	return err
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
func (m *moduleRepository) GetAllRoles() ([]models.RoleItems, error) {
	var roles []models.RoleItems
	rs := m.conn.Find(&roles)
	return roles, rs.Error
}
func (m *moduleRepository) SetPermission(moduleID, roleID uint, exec int) error {
	permission := models.Permissions{
		ModuleID:   moduleID,
		RoleItemID: roleID,
		IsExec:     exec,
	}
	rs := m.conn.Where("module_id = ? AND role_item_id = ?", moduleID, roleID).First(&permission)
	if rs.Error != nil {
		rs := m.conn.Create(&permission)
		return rs.Error
	}
	rs = m.conn.Model(permission).Omit("id").Where("id = ?", permission.ID).Updates(permission)
	return rs.Error
}

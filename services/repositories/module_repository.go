package repositories

import (
	"strconv"

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
func (m *moduleRepository) GetModule(module *models.Modules, id uint) error {
	err := m.conn.Where("ID = ?", id).First(module).Error
	return err
}
func (m *moduleRepository) CreateModule(module *models.Modules) error {
	if err := m.conn.Create(module).Error; err != nil {
		return err
	}
	return nil
}
func (m *moduleRepository) UpdateModule(module *models.Modules, id uint) error {
	if err := m.conn.Model(module).Omit("id").Where("id = ?", id).Updates(module).Error; err != nil {
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
func (m *moduleRepository) GetAllModules(modules *[]models.Modules, keyword, sorting, sortField, page, limit, focus string) (err error) {
	l, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return err
	}
	offset := (p * l) - l
	if focus == "inbox" {
		if keyword == "all" {
			if err := m.conn.Model(&models.Modules{}).Limit(l).Offset(offset).Order(sortField + " " + sorting).Find(&modules).Error; err != nil {
				return err
			}
			return nil
		}
		if err := m.conn.Model(&models.Modules{}).Where("modules.name like ?", "%"+keyword+"%").Limit(l).Offset(offset).Order(sortField + " " + sorting).Find(&modules).Error; err != nil {
			return err
		}
		return nil
	}
	if keyword == "all" {
		if err := m.conn.Unscoped().Model(&models.Modules{}).Where("modules.deleted_at IS NOT NULL").Limit(l).Offset(offset).Order(sortField + " " + sorting).Find(&modules).Error; err != nil {
			return err
		}
		return nil
	}
	if err := m.conn.Unscoped().Model(&models.Modules{}).Where("modules.deleted_at IS NOT NULL AND modules.name like ?", "%"+keyword+"%").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&modules).Error; err != nil {
		return err
	}
	return nil
}

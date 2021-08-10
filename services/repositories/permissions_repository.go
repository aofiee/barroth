package repositories

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"gorm.io/gorm"
)

type (
	PermissionsRepository struct {
		conn *gorm.DB
	}
)

func NewPermissionsRepository(conn *gorm.DB) domains.PermissionsRepository {
	return &PermissionsRepository{conn}
}
func (p *PermissionsRepository) SetPermissions(m *models.Permissions) error {
	err := p.conn.Model(models.Permissions{}).Omit("id").Where("module_id = ? AND role_item_id = ?", m.ModuleID, m.RoleItemID).Updates(m).Error
	return err
}

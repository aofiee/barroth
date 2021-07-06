package repositories

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"gorm.io/gorm"
)

type (
	roleRepository struct {
		conn *gorm.DB
	}
)

func NewRoleRepository(conn *gorm.DB) domains.RoleRepository {
	return &roleRepository{conn}
}
func (r *roleRepository) GetRole(role *models.RoleItems, id string) error {
	if err := r.conn.Where("id = ?", id).First(role).Error; err != nil {
		return err
	}
	return nil
}
func (r *roleRepository) CreateRole(role *models.RoleItems) error {
	if err := r.conn.Create(role).Error; err != nil {
		return err
	}
	return nil
}
func (r *roleRepository) UpdateRole(role *models.RoleItems, id string) error {
	if err := r.conn.Model(role).Omit("id").Where("id = ?").Updates(role).Error; err != nil {
		return err
	}
	return nil
}
func (r *roleRepository) GetAllRoles(roles *[]models.RoleItems) (err error) {
	if err := r.conn.Find(roles).Error; err != nil {
		return err
	}
	return nil
}

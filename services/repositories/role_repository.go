package repositories

import (
	"strconv"

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
func (r *roleRepository) GetAllRoles(roles *[]models.RoleItems, keyword, sorting, sortField, page, limit, focus string) (err error) {
	l, err := strconv.Atoi(limit)
	if err != nil {
		return err
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return err
	}
	p = p - 1
	if focus == "inbox" {
		if keyword == "all" {
			if err := r.conn.Model(&models.RoleItems{}).Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&roles).Error; err != nil {
				return err
			}
			return nil
		}
		if err := r.conn.Model(&models.RoleItems{}).Where("role_items.name like ?", "%"+keyword+"%").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&roles).Error; err != nil {
			return err
		}
		return nil
	}
	if keyword == "all" {
		if err := r.conn.Unscoped().Model(&models.RoleItems{}).Where("role_items.deleted_at IS NOT NULL").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&roles).Error; err != nil {
			return err
		}
		return nil
	}
	if err := r.conn.Unscoped().Model(&models.RoleItems{}).Where("role_items.deleted_at IS NOT NULL AND role_items.name like ?", "%"+keyword+"%").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&roles).Error; err != nil {
		return err
	}
	return nil
}
func (r *roleRepository) DeleteRoles(focus string, id []int) (int64, error) {
	if focus == "inbox" {
		rs := r.conn.Where("id IN ?", id).Delete(&models.RoleItems{})
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	rs := r.conn.Unscoped().Where("id IN ?", id).Delete(&models.RoleItems{})
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

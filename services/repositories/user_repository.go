package repositories

import (
	"strconv"

	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"gorm.io/gorm"
)

type (
	userRepository struct {
		conn *gorm.DB
	}
)

func NewUserRepository(conn *gorm.DB) domains.UserRepository {
	return &userRepository{conn}
}
func (u *userRepository) CreateUser(m *models.Users) error {
	if err := u.conn.Create(m).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) GetUser(m *models.Users, uuid string) error {
	if err := u.conn.Where("uuid = ?", uuid).First(m).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) GetAllUser(m *[]models.Users, keyword, sorting, sortField, page, limit, focus string) (err error) {
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
			if err := u.conn.Model(&models.Users{}).Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
				return err
			}
			return nil
		}
		if err := u.conn.Model(&models.Users{}).Where("users.name like ?", "%"+keyword+"%").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
			return err
		}
		return nil
	}
	if keyword == "all" {
		if err := u.conn.Unscoped().Model(&models.Users{}).Where("users.deleted_at IS NOT NULL").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
			return err
		}
		return nil
	}
	if err := u.conn.Unscoped().Model(&models.Users{}).Where("users.deleted_at IS NOT NULL AND users.name like ?", "%"+keyword+"%").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) UpdateUser(m *models.Users, uuid string) error {
	if err := u.conn.Model(m).Omit("id", "uuid").Where("uuid = ?", uuid).Updates(m).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) DeleteUsers(focus string, uuid []int) (int64, error) {
	if focus == "inbox" {
		rs := u.conn.Where("uuid IN ?", uuid).Delete(&models.Users{})
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	rs := u.conn.Unscoped().Where("uuid IN ?", uuid).Delete(&models.Users{})
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}
func (u *userRepository) RestoreUsers(id []int) (int64, error) {
	rs := u.conn.Unscoped().Model(&models.Users{}).Where("uuid IN ?", id).Update("deleted_at", nil)
	if rs.Error != nil {
		return 0, rs.Error
	}
	return rs.RowsAffected, nil
}

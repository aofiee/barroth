package repositories

import (
	"strconv"

	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"golang.org/x/crypto/bcrypt"
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
func (u *userRepository) GetUserByEmail(m *models.Users, email string) (err error) {
	if err := u.conn.Where("email = ?", email).First(m).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) GetAllUsers(m *[]models.Users, keyword, sorting, sortField, page, limit, focus string) (rows int64, err error) {
	l, err := strconv.Atoi(limit)
	if err != nil {
		return
	}
	p, err := strconv.Atoi(page)
	if err != nil {
		return
	}
	offset := (p * l) - l
	if focus == "inbox" {
		if keyword == "all" {
			if err = u.conn.Model(&models.Users{}).Limit(l).Offset(offset).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
				return
			}
			u.conn.Model(&models.Users{}).Count(&rows)
			return
		}
		if err = u.conn.Model(&models.Users{}).Where("users.name like ?", "%"+keyword+"%").Limit(l).Offset(offset).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
			return
		}
		u.conn.Model(&models.Users{}).Where("users.name like ?", "%"+keyword+"%").Count(&rows)
		return
	}
	if keyword == "all" {
		if err = u.conn.Unscoped().Model(&models.Users{}).Where("users.deleted_at IS NOT NULL").Limit(l).Offset(offset).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
			return
		}
		u.conn.Unscoped().Model(&models.Users{}).Where("users.deleted_at IS NOT NULL").Count(&rows)
		return
	}
	if err = u.conn.Unscoped().Model(&models.Users{}).Where("users.deleted_at IS NOT NULL AND users.name like ?", "%"+keyword+"%").Limit(l).Offset(p).Order(sortField + " " + sorting).Find(&m).Error; err != nil {
		return
	}
	u.conn.Unscoped().Model(&models.Users{}).Where("users.deleted_at IS NOT NULL AND users.name like ?", "%"+keyword+"%").Count(&rows)
	return
}
func (u *userRepository) UpdateUser(m *models.Users, uuid string) error {
	if err := u.conn.Model(m).Omit("id", "uuid").Where("uuid = ?", uuid).Updates(m).Error; err != nil {
		return err
	}
	return nil
}
func (u *userRepository) DeleteUsers(focus string, uuid []string) (int64, error) {
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
func (u *userRepository) HashPassword(user *models.Users) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(bytes)
	return err
}

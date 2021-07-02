package models

import (
	"gorm.io/gorm"
)

type (
	Users struct {
		gorm.Model
		Email      string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		Password   string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		Name       string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		Telephone  string    `gorm:"type:VARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		Image      string    `gorm:"type:VARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		UUID       string    `gorm:"type:VARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		UserRoleID UserRoles `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
		Status     int       `gorm:"type:TINYINT(1)"`
	}
	UserRoles struct {
		gorm.Model
		RoleItemID uint
		UserID     uint
	}
)

func (t *UserRoles) TableName() string {
	return "user_roles"
}
func (t *Users) TableName() string {
	return "users"
}

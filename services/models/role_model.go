package models

import (
	"gorm.io/gorm"
)

type (
	RoleItems struct {
		gorm.Model
		Name          string        `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		Description   string        `gorm:"type:TEXT CHARACTER SET utf8 COLLATE utf8_general_ci"`
		UserRoleID    []UserRoles   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoleItemID;references:ID"`
		PermissionsID []Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignKey:RoleItemID;references:ID"`
	}
)

func (t *RoleItems) TableName() string {
	return "role_items"
}

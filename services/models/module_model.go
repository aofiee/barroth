package models

import (
	"gorm.io/gorm"
)

type (
	Modules struct {
		gorm.Model
		Name          string        `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		Description   string        `gorm:"type:TEXT CHARACTER SET utf8 COLLATE utf8_general_ci"`
		ModuleSlug    string        `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		PermissionsID []Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignKey:ModuleID;references:ID"`
	}
)

func (t *Modules) TableName() string {
	return "modules"
}
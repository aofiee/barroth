package models

import (
	"gorm.io/gorm"
)

type (
	Permissions struct {
		gorm.Model
		ModuleID   uint
		RoleItemID uint
		IsExec     int `gorm:"type:TINYINT(1);default:0"`
	}
)

func (t *Permissions) TableName() string {
	return "permissions"
}

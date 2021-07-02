package models

import (
	"gorm.io/gorm"
)

type (
	Permissions struct {
		gorm.Model
		ModuleID   uint
		RoleItemID uint
		IsGet      int `gorm:"type:TINYINT(1);default:0"`
		IsPost     int `gorm:"type:TINYINT(1);default:0"`
		IsPut      int `gorm:"type:TINYINT(1);default:0"`
		IsDelete   int `gorm:"type:TINYINT(1);default:0"`
	}
)

func (t *Permissions) TableName() string {
	return "permissions"
}

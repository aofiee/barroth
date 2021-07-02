package models

import (
	"gorm.io/gorm"
)

type (
	System struct {
		gorm.Model
		AppName   string `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		SiteURL   string `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		IsInstall int    `gorm:"type:TINYINT(1);default:0"`
	}
)

func (t *System) TableName() string {
	return "system"
}

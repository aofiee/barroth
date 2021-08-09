package models

import (
	"time"

	"gorm.io/gorm"
)

type (
	BarrothModel struct {
		ID        uint           `gorm:"primaryKey" json:"id"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	}
	System struct {
		BarrothModel
		AppName   string `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		SiteURL   string `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		IsInstall int    `gorm:"type:TINYINT(1);default:0"`
	}
)

func (t *System) TableName() string {
	return "system"
}

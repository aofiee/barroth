package models

import (
	"time"
)

type (
	BarrothModel struct {
		ID        uint       `gorm:"primary_key" json:"id"`
		CreatedAt time.Time  `json:"created_at"`
		UpdatedAt time.Time  `json:"updated_at"`
		DeletedAt *time.Time `json:"deleted_at"`
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

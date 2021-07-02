package models

import (
	"gorm.io/gorm"
)

type (
	System struct {
		/*
			Id        uint64       `faker:"-"`
			CreatedAt time.Time    `gorm:"type:DATETIME NULL DEFAULT CURRENT_TIMESTAMP" faker:"-"`
			UpdatedAt time.Time    `gorm:"type:DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" faker:"-"`
			DeletedAt sql.NullTime `gorm:"type:DATETIME NULL DEFAULT NULL" faker:"-"`
			AppName   string       `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
			SiteURL   string       `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
			IsInstall int          `gorm:"type:TINYINT(1);default:0"`
		*/
		gorm.Model
		AppName   string `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		SiteURL   string `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci"`
		IsInstall int    `gorm:"type:TINYINT(1);default:0"`
	}
)

func (t *System) TableName() string {
	return "system"
}

package models

type (
	Modules struct {
		BarrothModel
		Name          string        `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"name"`
		Description   string        `gorm:"type:TEXT CHARACTER SET utf8 COLLATE utf8_general_ci" json:"description"`
		Method        string        `gorm:"type:VARCHAR(6) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"method"`
		ModuleSlug    string        `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"module_slug"`
		PermissionsID []Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignKey:ModuleID;references:ID" json:"-"`
	}
	ModuleMethodSlug struct {
		Name        string
		Description string
		Method      string
		Slug        string
	}
)

func (t *Modules) TableName() string {
	return "modules"
}

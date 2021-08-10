package models

type (
	Permissions struct {
		BarrothModel
		ModuleID   uint
		RoleItemID uint
		IsExec     *int `gorm:"type:TINYINT(1);default:0"`
	}
)

func (t *Permissions) TableName() string {
	return "permissions"
}

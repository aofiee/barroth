package models

type (
	RoleItems struct {
		BarrothModel
		Name          string        `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"name" validate:"required,min=6,max=255"`
		Description   string        `gorm:"type:TEXT CHARACTER SET utf8 COLLATE utf8_general_ci" json:"description" validate:"required,min=6"`
		UserRoleID    []UserRoles   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL; foreignKey:RoleItemID;references:ID" json:"-" validate:"-"`
		PermissionsID []Permissions `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignKey:RoleItemID;references:ID" json:"-" validate:"-"`
	}
)

func (t *RoleItems) TableName() string {
	return "role_items"
}

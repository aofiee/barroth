package models

type (
	Users struct {
		BarrothModel
		Email      string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"email"`
		Password   string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"-"`
		Name       string    `gorm:"type:VARCHAR(255) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"name"`
		Telephone  string    `gorm:"type:VARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"telephone"`
		Image      string    `gorm:"type:VARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"image"`
		UUID       string    `gorm:"type:VARCHAR(50) CHARACTER SET utf8 COLLATE utf8_general_ci" json:"uuid"`
		UserRoleID UserRoles `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE; foreignKey:UserID;references:ID" json:"-"`
		Status     int       `gorm:"type:TINYINT(1)" json:"status"`
	}
	UserRoles struct {
		BarrothModel
		RoleItemID uint
		UserID     uint
	}
) //foreignKey:LocationID;references:ID

func (t *UserRoles) TableName() string {
	return "user_roles"
}
func (t *Users) TableName() string {
	return "users"
}

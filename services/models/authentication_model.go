package models

type (
	Token struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
	TokenDetail struct {
		Token           Token        `json:"token"`
		AccessUUID      string       `json:"access_uuid"`
		RefreshUUID     string       `json:"refresh_uuid"`
		AccessTokenExp  int64        `json:"access_token_exp"`
		RefreshTokenExp int64        `json:"refresh_token_exp"`
		Context         TokenContext `json:"context"`
	}
	TokenContext struct {
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
		Role        string `json:"role"`
	}
	TokenRoleName struct {
		ID   string `json:"id" form:"id"`
		Name string `json:"name" form:"name"`
	}
)

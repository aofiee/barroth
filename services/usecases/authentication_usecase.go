package usecases

import (
	"errors"
	"strings"
	"time"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2/utils"
)

const (
	timeLoc = "Asia/Bangkok"
)

type (
	authenticationUseCase struct {
		authenticationRepo domains.AuthenticationRepository
	}
)

func NewAuthenticationUseCase(repo domains.AuthenticationRepository) domains.AuthenticationUseCase {
	return &authenticationUseCase{
		authenticationRepo: repo,
	}
}
func (a *authenticationUseCase) Login(m *models.Users, email, password string) error {
	err := a.authenticationRepo.Login(m, email)
	if err != nil {
		return err
	}
	ok := a.authenticationRepo.CheckPasswordHash(m, password)
	if !ok {
		return errors.New(constants.ERR_USERNAME_PASSWORD_INCORRECT)
	}
	return nil
}
func (a *authenticationUseCase) CreateToken(m *models.Users) (models.TokenDetail, error) {
	var token models.TokenDetail
	location, _ := time.LoadLocation(timeLoc)

	accessToken := time.Now().In(location).Add(time.Minute * 15).Unix()
	refreshToken := time.Now().In(location).Add(time.Hour * 24 * 7).Unix()

	token.AccessTokenExp = accessToken
	token.RefreshTokenExp = refreshToken
	token.AccessUUID = utils.UUIDv4()
	token.RefreshUUID = utils.UUIDv4()

	context := models.TokenContext{
		Email:       m.Email,
		DisplayName: m.Name,
	}
	var role models.TokenRoleName
	err := a.authenticationRepo.GetRoleNameByUserID(&role, m.ID)
	if err != nil {
		return models.TokenDetail{}, err
	}
	context.Role = role.Name
	token.Context = context

	return token, nil
}

func (a *authenticationUseCase) GenerateAccessTokenBy(u *models.Users, t *models.TokenDetail) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	location, _ := time.LoadLocation(timeLoc)
	claims["iss"] = barroth_config.ENV.AppName
	claims["sub"] = u.UUID
	claims["exp"] = t.AccessTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["context"] = t.Context
	claims["access_uuid"] = t.AccessUUID
	rs, err := token.SignedString([]byte(barroth_config.ENV.AccessKey))
	t.Token.AccessToken = rs
	return err
}
func (a *authenticationUseCase) GenerateRefreshTokenBy(u *models.Users, t *models.TokenDetail) error {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	location, _ := time.LoadLocation(timeLoc)
	claims["iss"] = barroth_config.ENV.AppName
	claims["sub"] = u.UUID
	claims["exp"] = t.RefreshTokenExp
	claims["iat"] = time.Now().In(location).Unix()
	claims["refresh_uuid"] = t.RefreshUUID
	rs, err := token.SignedString([]byte(barroth_config.ENV.RefreshKey))
	t.Token.RefreshToken = rs
	return err
}
func (a *authenticationUseCase) SaveToken(uuid string, t *models.TokenDetail) error {
	loc, _ := time.LoadLocation(timeLoc)
	accessTokenExpire := time.Unix(t.AccessTokenExp, 0).In(loc)
	refreshTokenExpire := time.Unix(t.RefreshTokenExp, 0).In(loc)
	now := time.Now().In(loc)
	err := a.authenticationRepo.SaveToken(uuid, t.AccessUUID, accessTokenExpire.Sub(now))
	if err != nil {
		return err
	}
	err = a.authenticationRepo.SaveToken(uuid, t.RefreshUUID, refreshTokenExpire.Sub(now))
	if err != nil {
		return err
	}
	return nil
}
func (a *authenticationUseCase) DeleteToken(uuid string) error {
	err := a.authenticationRepo.DeleteToken(uuid)
	return err
}
func (a *authenticationUseCase) GetUser(u *models.Users, uuid string) error {
	err := a.authenticationRepo.GetUser(u, uuid)
	return err
}
func (a *authenticationUseCase) GetAccessUUIDFromRedis(uuid string) (string, error) {
	result, err := a.authenticationRepo.GetAccessUUIDFromRedis(uuid)
	return result, err
}
func (a *authenticationUseCase) CheckRoutePermission(roleName, method, slug string) bool {
	folder := strings.Split(slug, "/")
	checkRouting := ""
	for _, v := range folder {
		if v != "" {
			checkRouting += "/" + v
			ok := a.authenticationRepo.CheckRoutePermission(roleName, method, checkRouting)
			if ok {
				return true
			}
		}
	}
	return false
}

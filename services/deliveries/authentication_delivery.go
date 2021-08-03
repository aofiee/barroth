package deliveries

import (
	"fmt"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt"
)

type (
	authenticationHandler struct {
		authenticationUseCase domains.AuthenticationUseCase
		moduleName            string
		description           string
	}
	paramsLogin struct {
		Email    string `json:"email" validate:"required,email,min=6,max=255"`
		Password string `json:"password" validate:"required,min=6,max=64"`
	}
)

func NewAuthenHandler(usecase domains.AuthenticationUseCase, m, d string, u *[]models.ModuleMethodSlug) *authenticationHandler {
	moduleRepo := repositories.NewModuleRepository(databases.DB)
	moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
	for _, value := range *u {
		newModule := models.Modules{
			Name:        m,
			Description: d,
			ModuleSlug:  value.Slug,
			Method:      value.Method,
		}
		err := moduleUseCase.GetModuleBySlug(&newModule, value.Method, value.Slug)
		if err != nil {
			moduleUseCase.CreateModule(&newModule)
		}
	}
	return &authenticationHandler{
		authenticationUseCase: usecase,
		moduleName:            m,
		description:           d,
	}
}

func GetAuthHandlerUsecase(usecase domains.AuthenticationUseCase) *authenticationHandler {
	return &authenticationHandler{
		authenticationUseCase: usecase,
	}
}

func (a *authenticationHandler) Login(c *fiber.Ctx) error {
	var auth paramsLogin
	err := c.BodyParser(&auth)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&auth)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	var u models.Users
	err = a.authenticationUseCase.Login(&u, auth.Email, auth.Password)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_USERNAME_PASSWORD_INCORRECT, fiber.StatusUnauthorized)
	}
	token, err := a.authenticationUseCase.CreateToken(&u)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_GET_ROLE_NAME, fiber.StatusBadRequest)
	}
	err = a.authenticationUseCase.GenerateAccessTokenBy(&u, &token)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_TOKEN_CANNOT_SIGNED_KEY, fiber.StatusBadRequest)
	}
	err = a.authenticationUseCase.GenerateRefreshTokenBy(&u, &token)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_TOKEN_CANNOT_SIGNED_KEY, fiber.StatusBadRequest)
	}
	err = a.authenticationUseCase.SaveToken(u.UUID, &token)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_SAVE_TOKEN_TO_REDIS, fiber.StatusInternalServerError)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_LOGIN_SUCCESSFUL,
		"error": nil,
		"data":  token.Token,
	})
}
func (a *authenticationHandler) Logout(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	accessUUID := claims["access_uuid"].(string)
	err := a.authenticationUseCase.DeleteToken(accessUUID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_DELETE_TOKEN_TO_REDIS, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_LOGOUT_COMPLETED,
		"error": nil,
	})
}
func (a *authenticationHandler) AuthorizationRequired() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SuccessHandler: a.AuthSuccess,
		ErrorHandler:   a.AuthError,
		SigningKey:     []byte(barroth_config.ENV.AccessKey),
		SigningMethod:  "HS256",
		TokenLookup:    "header:Authorization",
		AuthScheme:     "Bearer",
	})
}
func (a *authenticationHandler) AuthError(c *fiber.Ctx, e error) error {
	return helpers.FailOnError(c, e, "Unauthorized", fiber.StatusUnauthorized)
}

func (a *authenticationHandler) AuthSuccess(c *fiber.Ctx) error {
	return c.Next()
}
func (a *authenticationHandler) RefreshToken(c *fiber.Ctx) error {
	var param models.RefreshToken
	err := c.BodyParser(&param)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	token, err := jwt.Parse(param.Token, a.VerifyToken)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_TOKEN_SIGNED_NOT_MATCH, fiber.StatusUnauthorized)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uuid, ok := claims["sub"].(string)
		if !ok {
			return helpers.FailOnError(c, err, constants.ERR_TOKEN_CANNOT_SIGNED_KEY, fiber.StatusNotFound)
		}
		refreshAccesss, ok := claims["refresh_uuid"].(string)
		if !ok {
			return helpers.FailOnError(c, err, constants.ERR_TOKEN_CANNOT_SIGNED_KEY, fiber.StatusNotFound)
		}
		err = a.authenticationUseCase.DeleteToken(refreshAccesss)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_DELETE_TOKEN_TO_REDIS, fiber.StatusBadRequest)
		}
		var user models.Users
		err := a.authenticationUseCase.GetUser(&user, uuid)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_GET_USER_BY_UUID_NOT_FOUND, fiber.StatusNotFound)
		}
		renewToken, err := a.authenticationUseCase.CreateToken(&user)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_GET_ROLE_NAME, fiber.StatusBadRequest)
		}
		err = a.authenticationUseCase.GenerateAccessTokenBy(&user, &renewToken)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_TOKEN_CANNOT_SIGNED_KEY, fiber.StatusBadRequest)
		}
		err = a.authenticationUseCase.GenerateRefreshTokenBy(&user, &renewToken)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_TOKEN_CANNOT_SIGNED_KEY, fiber.StatusBadRequest)
		}
		err = a.authenticationUseCase.SaveToken(user.UUID, &renewToken)
		if err != nil {
			return helpers.FailOnError(c, err, constants.ERR_CANNOT_SAVE_TOKEN_TO_REDIS, fiber.StatusInternalServerError)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"msg":   constants.ERR_REFRESH_TOKEN_SUCCESSFUL,
			"error": nil,
			"data":  renewToken.Token,
		})
	} else {
		return helpers.FailOnError(c, err, constants.ERR_REFRESH_TOKEN_EXPIRE, fiber.StatusUnauthorized)
	}
}
func (a *authenticationHandler) VerifyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte(barroth_config.ENV.RefreshKey), nil
}
func (a *authenticationHandler) IsRevokeToken(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	accessUUID := claims["access_uuid"].(string)
	_, err := a.authenticationUseCase.GetAccessUUIDFromRedis(accessUUID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_REFRESH_TOKEN_EXPIRE, fiber.StatusUnauthorized)
	}
	return c.Next()
}

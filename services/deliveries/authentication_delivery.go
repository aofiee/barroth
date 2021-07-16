package deliveries

import (
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
		slug                  string
	}
	paramsLogin struct {
		Email    string `json:"email" validate:"required,email,min=6,max=255"`
		Password string `json:"password" validate:"required,min=6,max=64"`
	}
)

func NewAuthenHandler(usecase domains.AuthenticationUseCase, m, d, u string) *authenticationHandler {
	newModule := models.Modules{
		Name:        m,
		Description: d,
		ModuleSlug:  u,
	}
	moduleRepo := repositories.NewModuleRepository(databases.DB)
	moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
	err := moduleUseCase.GetModule(&newModule, u)
	if err != nil {
		moduleUseCase.CreateModule(&newModule)
	}
	return &authenticationHandler{
		authenticationUseCase: usecase,
		moduleName:            m,
		description:           d,
		slug:                  u,
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
		return helpers.FailOnError(c, err, "StatusUnauthorized", fiber.StatusUnauthorized)
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

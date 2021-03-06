package deliveries

import (
	"strconv"
	"strings"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type (
	userHandler struct {
		userUseCase domains.UserUseCase
	}
	paramsUser struct {
		Email     string `json:"email" form:"email" validate:"required,email,min=6,max=255"`
		Password  string `json:"password" form:"password" validate:"required,min=6,max=64"`
		Telephone string `json:"telephone" form:"telephone" validate:"required,min=10,max=50"`
		Name      string `json:"name" form:"name" validate:"required,min=6,max=255"`
		RoleID    int    `json:"role_id" form:"role_id" validate:"required,number"`
	}
	paramsRegister struct {
		Email       string `json:"email" form:"email" validate:"required,email,min=6,max=255"`
		UUID        string `json:"uuid" form:"uuid" validate:"required,min=28,max=128"`
		Provider    string `json:"provider" form:"provider" validate:"required,min=5,max=255"`
		DisplayName string `json:"display_name" form:"display_name"`
		PhotoURL    string `json:"photo_url" form:"photo_url"`
	}
	paramUUID struct {
		UsersID []string `json:"users_id" validate:"required"`
	}
	// paramsGetAllUsers struct {
	// 	Keyword   string `json:"keyword" form:"keyword"`
	// 	Page      string `json:"page" form:"page"`
	// 	Limit     string `json:"limit" form:"limit"`
	// 	Sorting   string `json:"sort" form:"sort" validate:"eq=desc|eq=asc"`
	// 	SortField string `json:"field" form:"field" validate:"eq=id|eq=name|eq=email|eq=password|eq=telephone|eq=uuid|eq=user_role_id|eq=status"`
	// 	Focus     string `json:"focus" form:"focus" validate:"eq=inbox|eq=trash"`
	// }
)

const (
	effectedConst = " effected "
	itemsConst    = " items"
)

func NewUserHandelr(usecase domains.UserUseCase, u *[]models.ModuleMethodSlug) *userHandler {
	moduleRepo := repositories.NewModuleRepository(databases.DB)
	moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
	for _, value := range *u {
		newModule := models.Modules{
			Name:        value.Name,
			Description: value.Description,
			ModuleSlug:  value.Slug,
			Method:      value.Method,
		}
		err := moduleUseCase.GetModuleBySlug(&newModule, value.Method, value.Slug)
		if err != nil {
			moduleUseCase.CreateModule(&newModule)
		}
	}
	return &userHandler{
		userUseCase: usecase,
	}
}

func (u *userHandler) RegisterUser(c *fiber.Ctx) error {
	var nu paramsRegister
	err := c.BodyParser(&nu)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&nu)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	user := models.Users{
		Email:    nu.Email,
		Name:     nu.DisplayName,
		UUID:     nu.UUID,
		Provider: nu.Provider,
		Image:    nu.PhotoURL,
	}
	err = u.userUseCase.CreateUser(&user)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_USER, fiber.StatusBadRequest)
	}
	roldID, _ := strconv.ParseUint(barroth_config.ENV.CustomerRoleID, 10, 32)
	role := models.UserRoles{
		RoleItemID: uint(roldID),
		UserID:     user.ID,
	}
	err = u.userUseCase.SetUserRole(&role, user.ID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_CREATE_USER_SUCCESSFUL,
		"error": nil,
	})
}

func (u *userHandler) NewUser(c *fiber.Ctx) error {
	var nu paramsUser
	err := c.BodyParser(&nu)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&nu)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	user := models.Users{
		Email:     nu.Email,
		Password:  nu.Password,
		Name:      nu.Name,
		Telephone: nu.Telephone,
		Provider:  "backend",
	}
	err = u.userUseCase.CreateUser(&user)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_USER, fiber.StatusBadRequest)
	}
	role := models.UserRoles{
		RoleItemID: uint(nu.RoleID),
		UserID:     user.ID,
	}
	err = u.userUseCase.SetUserRole(&role, user.ID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_CREATE_USER_SUCCESSFUL,
		"error": nil,
	})
}
func (u *userHandler) UpdateUser(c *fiber.Ctx) error {
	var nu paramsUser
	uuid := c.Params("id")
	err := c.BodyParser(&nu)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&nu)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	user := models.Users{
		Email:     nu.Email,
		Password:  nu.Password,
		Name:      nu.Name,
		Telephone: nu.Telephone,
	}
	err = u.userUseCase.UpdateUser(&user, uuid)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_UPDATE_USER, fiber.StatusBadRequest)
	}
	err = u.userUseCase.GetUser(&user, uuid)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_GET_USER_FAIL, fiber.StatusBadRequest)
	}
	role := models.UserRoles{
		RoleItemID: uint(nu.RoleID),
		UserID:     user.ID,
	}
	err = u.userUseCase.SetUserRole(&role, user.ID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_UPDATED_USER_SUCCESSFUL,
		"error": nil,
	})
}
func (u *userHandler) GetUser(c *fiber.Ctx) error {
	param := c.Params("id")
	if param == "me" {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		param = claims["sub"].(string)
	}
	var user models.Users
	err := u.userUseCase.GetUser(&user, param)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_GET_USER_FAIL, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_GET_USER_SUCCESSFUL,
		"error": nil,
		"data":  user,
	})
}
func (u *userHandler) DeleteUser(c *fiber.Ctx) error {
	param := c.Params("id")
	uuids := []string{param}
	effectRows, err := u.userUseCase.DeleteUsers("inbox", uuids)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_DELETE_USER_SUCCESSFUL, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_DELETE_USER_SUCCESSFUL + effectedConst + strconv.FormatInt(effectRows, 10) + itemsConst,
		"error": nil,
	})
}
func (u *userHandler) DeleteMultitpleUsers(c *fiber.Ctx) error {
	var param paramUUID
	focus := c.Query("focus")
	if focus != "inbox" && focus != "trash" {
		focus = "inbox"
	}
	err := c.BodyParser(&param)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&param)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	effectRows, err := u.userUseCase.DeleteUsers(focus, param.UsersID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_DELETE_USER_SUCCESSFUL, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_DELETE_USER_SUCCESSFUL + effectedConst + strconv.FormatInt(effectRows, 10) + itemsConst,
		"error": nil,
	})
}
func (r *userHandler) BuildGetAllUsersParam(keyword, page, limit, sort, field, focus string) paramsGetAllRoles {
	if keyword == "" {
		keyword = "all"
	}
	if page == "" {
		page = "0"
	}
	if limit == "" {
		limit = "10"
	}
	if sort == "" {
		sort = "asc"
	}
	if field == "" {
		field = "id"
	}
	if focus == "" {
		focus = "inbox"
	}
	var param paramsGetAllRoles
	param.SortField = field
	param.Keyword = keyword
	param.Focus = focus
	param.Limit = limit
	param.Sorting = sort
	param.Page = page
	return param
}
func (u *userHandler) GetAllUsers(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	page := c.Query("page")
	limit := c.Query("limit")
	sort := strings.ToLower(c.Query("sort"))
	field := strings.ToLower(c.Query("field"))
	focus := strings.ToLower(c.Query("focus"))
	param := u.BuildGetAllUsersParam(keyword, page, limit, sort, field, focus)
	errorResponse := helpers.ValidateStruct(&param)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	var users []models.Users
	rows, err := u.userUseCase.GetAllUsers(&users, param.Keyword, param.Sorting, param.SortField, param.Page, param.Limit, param.Focus)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_GET_ALL_ROLES, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":     constants.ERR_GET_USER_SUCCESSFUL,
		"error":   nil,
		"data":    users,
		"total":   rows,
		"current": param.Page,
	})
}
func (u *userHandler) RestoreUsers(c *fiber.Ctx) error {
	var param paramUUID
	err := c.BodyParser(&param)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&param)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	effectRows, err := u.userUseCase.RestoreUsers(param.UsersID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_RESTORE_USER_FROM_TRASH_TO_INBOX_SUCCESSFUL, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_RESTORE_USER_FROM_TRASH_TO_INBOX_SUCCESSFUL + effectedConst + strconv.FormatInt(effectRows, 10) + itemsConst,
		"error": nil,
	})
}

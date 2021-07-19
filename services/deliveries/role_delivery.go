package deliveries

import (
	"strconv"
	"strings"

	"github.com/aofiee/barroth/constants"
	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/helpers"
	"github.com/aofiee/barroth/models"
	"github.com/aofiee/barroth/repositories"
	"github.com/aofiee/barroth/usecases"
	"github.com/gofiber/fiber/v2"
)

type (
	roleHandler struct {
		roleUseCase domains.RoleUseCase
		moduleName  string
		description string
	}
	paramsGetAllRoles struct {
		Keyword   string `json:"keyword" form:"keyword"`
		Page      string `json:"page" form:"page"`
		Limit     string `json:"limit" form:"limit"`
		Sorting   string `json:"sort" form:"sort" validate:"eq=desc|eq=asc"`
		SortField string `json:"field" form:"field" validate:"eq=id|eq=name|eq=email|eq=password|eq=telephone|eq=uuid|eq=user_role_id|eq=status"`
		Focus     string `json:"focus" form:"focus" validate:"eq=inbox|eq=trash"`
	}
	paramRequestRolesID struct {
		RoleID []int `json:"role_id" validate:"required"`
	}
)

func NewRoleHandelr(usecase domains.RoleUseCase, m, d string, u *[]models.ModuleMethodSlug) *roleHandler {
	for _, value := range *u {
		newModule := models.Modules{
			Name:        m,
			Description: d,
			ModuleSlug:  value.Slug,
			Method:      value.Method,
		}
		moduleRepo := repositories.NewModuleRepository(databases.DB)
		moduleUseCase := usecases.NewModuleUseCase(moduleRepo)
		err := moduleUseCase.GetModuleBySlug(&newModule, value.Method, value.Slug)
		if err != nil {
			moduleUseCase.CreateModule(&newModule)
		}
	}

	return &roleHandler{
		roleUseCase: usecase,
		moduleName:  m,
		description: d,
	}
}
func (r *roleHandler) NewRole(c *fiber.Ctx) error {
	var role models.RoleItems
	err := c.BodyParser(&role)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&role)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	err = r.roleUseCase.CreateRole(&role)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_CREATE_ROLE, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_CREATE_ROLE_SUCCESSFUL,
		"error": nil,
	})
}
func (r *roleHandler) BuildGetAllRolesParam(k, p, l, s, f, fo string) paramsGetAllRoles {
	keyword := k
	page := p
	limit := l
	sort := s
	field := f
	focus := fo
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
func (r *roleHandler) GetAllRoles(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	page := c.Query("page")
	limit := c.Query("limit")
	sort := strings.ToLower(c.Query("sort"))
	field := strings.ToLower(c.Query("field"))
	focus := strings.ToLower(c.Query("focus"))
	param := r.BuildGetAllRolesParam(keyword, page, limit, sort, field, focus)
	errorResponse := helpers.ValidateStruct(&param)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	var roles []models.RoleItems
	err := r.roleUseCase.GetAllRoles(&roles, param.Keyword, param.Sorting, param.SortField, param.Page, param.Limit, param.Focus)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_GET_ALL_ROLES, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_GET_ALL_ROLE_SUCCESSFULE,
		"error": nil,
		"data":  roles,
	})
}
func (r *roleHandler) DeleteRoles(c *fiber.Ctx) error {
	focus := strings.ToLower(c.Query("focus"))
	if focus == "" {
		focus = "inbox"
	}
	var params paramRequestRolesID
	err := c.BodyParser(&params)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errs := helpers.ValidateStruct(&params)
	if errs != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errs,
		})
	}
	rs, err := r.roleUseCase.DeleteRoles(focus, params.RoleID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_DELETE_ROLE, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "deleted " + strconv.FormatInt(rs, 10) + " roles successful.",
		"error": nil,
	})
}
func (r *roleHandler) RestoreRoles(c *fiber.Ctx) error {
	var params paramRequestRolesID
	err := c.BodyParser(&params)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errs := helpers.ValidateStruct(&params)
	if errs != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errs,
		})
	}
	rs, err := r.roleUseCase.RestoreRoles(params.RoleID)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_RESTORE_ROLES, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "restore " + strconv.FormatInt(rs, 10) + " roles successful.",
		"error": nil,
	})
}
func (r *roleHandler) GetRole(c *fiber.Ctx) error {
	var role models.RoleItems
	err := r.roleUseCase.GetRole(&role, c.Params("id"))
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_GET_ROLE_ID+c.Params("id"), fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "get data of role id " + c.Params("id") + " is completed.",
		"error": nil,
		"data":  role,
	})
}
func (r *roleHandler) UpdateRole(c *fiber.Ctx) error {
	var role models.RoleItems
	err := c.BodyParser(&role)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&role)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	err = r.roleUseCase.UpdateRole(&role, c.Params("id"))
	if err != nil {
		return helpers.FailOnError(c, err, "cannot update role id "+c.Params("id"), fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "update role id " + c.Params("id") + " is completed.",
		"error": nil,
	})
}

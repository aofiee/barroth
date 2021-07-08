package deliveries

import (
	"encoding/json"
	"strconv"
	"strings"

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
		slug        string
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

func NewRoleHandelr(usecase domains.RoleUseCase, m, d, u string) *roleHandler {
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
	return &roleHandler{
		roleUseCase: usecase,
		moduleName:  m,
		description: d,
		slug:        u,
	}
}
func (r *roleHandler) NewRole(c *fiber.Ctx) error {
	var role models.RoleItems
	err := c.BodyParser(&role)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot parse json", fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&role)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   "input error.",
			"error": errorResponse,
		})
	}
	err = r.roleUseCase.CreateRole(&role)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot create role", fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "create role successful.",
		"error": nil,
	})
}
func (r *roleHandler) BuildGetAllRolesParam(c *fiber.Ctx) ([]byte, error) {
	keyword := c.Query("keyword")
	page := c.Query("page")
	limit := c.Query("limit")
	sort := strings.ToLower(c.Query("sort"))
	field := strings.ToLower(c.Query("field"))
	focus := strings.ToLower(c.Query("focus"))
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
	json, err := json.Marshal(&param)
	if err != nil {
		return nil, err
	}
	return json, nil
}
func (r *roleHandler) GetAllRoles(c *fiber.Ctx) error {
	p, err := r.BuildGetAllRolesParam(c)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot parse params", fiber.StatusNotAcceptable)
	}
	var param paramsGetAllRoles
	json.Unmarshal(p, &param)
	errorResponse := helpers.ValidateStruct(&param)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   "input error.",
			"error": errorResponse,
		})
	}
	var roles []models.RoleItems
	err = r.roleUseCase.GetAllRoles(&roles, param.Keyword, param.Sorting, param.SortField, param.Page, param.Limit, param.Focus)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot get all roles", fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "get all role successful.",
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
		return helpers.FailOnError(c, err, "cannot parse json", fiber.StatusBadRequest)
	}
	errs := helpers.ValidateStruct(&params)
	if errs != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   "input error.",
			"error": errs,
		})
	}
	rs, err := r.roleUseCase.DeleteRoles(focus, params.RoleID)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot delete roles", fiber.StatusBadRequest)
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
		return helpers.FailOnError(c, err, "cannot parse json", fiber.StatusBadRequest)
	}
	errs := helpers.ValidateStruct(&params)
	if errs != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   "input error.",
			"error": errs,
		})
	}
	rs, err := r.roleUseCase.RestoreRoles(params.RoleID)
	if err != nil {
		return helpers.FailOnError(c, err, "cannot restore roles", fiber.StatusBadRequest)
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
		return helpers.FailOnError(c, err, "cannot get role id "+c.Params("id"), fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   "get data of role id " + c.Params("id") + " is completed.",
		"error": nil,
		"data":  role,
	})
}

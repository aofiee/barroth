package deliveries

import (
	"encoding/json"
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
	paramGetAllRoles struct {
		Keyword   string `json:"keyword" form:"keyword"`
		Page      string `json:"page" form:"page"`
		Limit     string `json:"limit" form:"limit"`
		Sorting   string `json:"sort" form:"sort" validate:"eq=desc|eq=asc"`
		SortField string `json:"field" form:"field" validate:"eq=id|eq=name|eq=email|eq=password|eq=telephone|eq=uuid|eq=user_role_id|eq=status"`
		Focus     string `json:"focus" form:"focus" validate:"eq=inbox|eq=trash"`
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
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
	var param paramGetAllRoles
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
	var param paramGetAllRoles
	json.Unmarshal(p, &param)
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

package deliveries

import (
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
	moduleHandler struct {
		moduleUseCase domains.ModuleUseCase
	}
	paramsGetAllModules struct {
		Keyword   string `json:"keyword" form:"keyword"`
		Page      string `json:"page" form:"page"`
		Limit     string `json:"limit" form:"limit"`
		Sorting   string `json:"sort" form:"sort" validate:"eq=desc|eq=asc"`
		SortField string `json:"field" form:"field" validate:"eq=id|eq=name|eq=email|eq=password|eq=telephone|eq=uuid|eq=user_role_id|eq=status"`
		Focus     string `json:"focus" form:"focus" validate:"eq=inbox|eq=trash"`
	}
	paramsModules struct {
		Name        string `json:"name" validate:"required"`
		Description string `json:"description" validate:"required"`
		Method      string `json:"method" validate:"required,eq=GET|eq=POST|eq=PUT|eq=DELETE|eq=OPTIONS"`
		ModuleSlug  string `json:"module_slug" validate:"required"`
	}
)

func NewModuleHandler(usecase domains.ModuleUseCase, u *[]models.ModuleMethodSlug) *moduleHandler {
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
	return &moduleHandler{
		moduleUseCase: usecase,
	}
}
func (r *moduleHandler) BuildGetAllRolesParam(keyword, page, limit, sort, field, focus string) paramsGetAllModules {
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
	var param paramsGetAllModules
	param.SortField = field
	param.Keyword = keyword
	param.Focus = focus
	param.Limit = limit
	param.Sorting = sort
	param.Page = page
	return param
}
func (m *moduleHandler) GetAllModules(c *fiber.Ctx) error {
	keyword := c.Query("keyword")
	page := c.Query("page")
	limit := c.Query("limit")
	sort := strings.ToLower(c.Query("sort"))
	field := strings.ToLower(c.Query("field"))
	focus := strings.ToLower(c.Query("focus"))
	param := m.BuildGetAllRolesParam(keyword, page, limit, sort, field, focus)
	errorResponse := helpers.ValidateStruct(&param)
	if errorResponse != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	var modules []models.Modules
	err := m.moduleUseCase.GetAllModules(&modules, param.Keyword, param.Sorting, param.SortField, param.Page, param.Limit, param.Focus)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_CANNOT_GET_ALL_ROLES, fiber.StatusBadRequest)
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_GET_ALL_ROLE_SUCCESSFULE,
		"error": nil,
		"data":  modules,
	})
}
func (m *moduleHandler) UpdateModule(c *fiber.Ctx) error {
	var param paramsModules
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
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_INPUT_ERROR,
			"error": errorResponse,
		})
	}
	var module models.Modules
	module.Name = param.Name
	module.Description = param.Description
	module.Method = param.Method
	module.ModuleSlug = param.ModuleSlug
	err = m.moduleUseCase.UpdateModule(&module, uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_CANNOT_UPDATE_MODULE,
			"error": errorResponse,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_UPDATE_MODULE_SUCCESSFUL,
		"error": nil,
	})
}

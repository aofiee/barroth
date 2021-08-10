package deliveries

import (
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
	permissionsHandler struct {
		permissionsUseCase domains.PermissionsUseCase
	}
	paramsSetPermission struct {
		Permissions []permission `json:"permissions" validate:"required"`
	}
	permission struct {
		ModuleID int `json:"module_id" validate:"number,required"`
		IsExec   int `json:"is_exec" validate:"number,required,max=1"`
	}
)

func NewPermissionHandler(usecase domains.PermissionsUseCase, u *[]models.ModuleMethodSlug) *permissionsHandler {
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
	return &permissionsHandler{
		permissionsUseCase: usecase,
	}
}
func (p *permissionsHandler) SetPermission(c *fiber.Ctx) error {
	var params paramsSetPermission
	err := c.BodyParser(&params)
	if err != nil {
		return helpers.FailOnError(c, err, constants.ERR_PARSE_JSON_FAIL, fiber.StatusBadRequest)
	}
	errorResponse := helpers.ValidateStruct(&params)
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
			"error": err,
		})
	}
	var permissions []models.Permissions
	for _, v := range params.Permissions {
		v := v
		permission := models.Permissions{
			ModuleID:   uint(v.ModuleID),
			RoleItemID: uint(id),
			IsExec:     &v.IsExec,
		}
		permissions = append(permissions, permission)
	}
	err = p.permissionsUseCase.SetPermissions(&permissions)
	if err != nil {
		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{
			"msg":   constants.ERR_CANNOT_SET_PERMISSIONS,
			"error": errorResponse,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":   constants.ERR_UPDATE_MODULE_SUCCESSFUL,
		"error": nil,
	})
}

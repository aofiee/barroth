package domains

import (
	"github.com/aofiee/barroth/models"
)

type (
	ModuleUseCase interface {
		CreateModule(m *models.Modules) (err error)
		GetModuleBySlug(m *models.Modules, method, slug string) (err error)
		GetModule(m *models.Modules, id uint) (err error)
		UpdateModule(m *models.Modules, id uint) (err error)
		GetAllRoles() (m []models.RoleItems, err error)
		SetPermission(moduleID, roleID uint, exec int) (err error)
		GetAllModules(m *[]models.Modules, keyword, sorting, sortField, page, limit, focus string) (err error)
	}
	ModuleRepository interface {
		CreateModule(m *models.Modules) (err error)
		GetModuleBySlug(m *models.Modules, method, slug string) (err error)
		GetModule(m *models.Modules, id uint) (err error)
		UpdateModule(m *models.Modules, id uint) (err error)
		GetAllRoles() (m []models.RoleItems, err error)
		SetPermission(moduleID, roleID uint, exec int) (err error)
		GetAllModules(m *[]models.Modules, keyword, sorting, sortField, page, limit, focus string) (err error)
	}
)

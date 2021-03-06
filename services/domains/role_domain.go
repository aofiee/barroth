package domains

import (
	"github.com/aofiee/barroth/models"
)

type (
	RoleUseCase interface {
		CreateRole(m *models.RoleItems) (err error)
		GetRole(m *models.RoleItems, id string) (err error)
		GetAllRoles(m *[]models.RoleItems, keyword, sorting, sortField, page, limit, focus string) (err error)
		UpdateRole(m *models.RoleItems, id string) (err error)
		DeleteRoles(focus string, id []int) (rs int64, err error)
		RestoreRoles(id []int) (rs int64, err error)
		GetAllModules() (m []models.Modules, err error)
		SetPermission(moduleID, roleID uint, exec int) (err error)
	}
	RoleRepository interface {
		CreateRole(m *models.RoleItems) (err error)
		GetRole(m *models.RoleItems, id string) (err error)
		GetAllRoles(m *[]models.RoleItems, keyword, sorting, sortField, page, limit, focus string) (err error)
		UpdateRole(m *models.RoleItems, id string) (err error)
		DeleteRoles(focus string, id []int) (rs int64, err error)
		RestoreRoles(id []int) (rs int64, err error)
		GetAllModules() (m []models.Modules, err error)
		SetPermission(moduleID, roleID uint, exec int) (err error)
	}
)

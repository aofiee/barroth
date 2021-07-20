package domains

import (
	"github.com/aofiee/barroth/models"
)

type (
	SystemUseCase interface {
		GetFirstSystemInstallation(s *models.System) (err error)
		CreateSystem(s *models.System) (err error)
		GetSystem(s *models.System, id string) (err error)
		UpdateSystem(s *models.System, id string) (err error)
		CreateUser(m *models.Users) (err error)
		CreateRole(m *models.RoleItems) (err error)
		SetExecToAllModules(m *[]models.Modules, roleID uint, isExec int) (err error)
	}
	SystemRepository interface {
		GetFirstSystemInstallation(s *models.System) (err error)
		CreateSystem(s *models.System) (err error)
		GetSystem(s *models.System, id string) (err error)
		UpdateSystem(s *models.System, id string) (err error)
		CreateUser(m *models.Users) (err error)
		CreateRole(m *models.RoleItems) (err error)
		HashPassword(user *models.Users) (err error)
		GetAllModules(m *[]models.Modules) (err error)
		SetPermissions(m *models.Permissions) (err error)
	}
)

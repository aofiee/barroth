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
	}
	SystemRepository interface {
		GetFirstSystemInstallation(s *models.System) (err error)
		CreateSystem(s *models.System) (err error)
		GetSystem(s *models.System, id string) (err error)
		UpdateSystem(s *models.System, id string) (err error)
		CreateUser(m *models.Users) (err error)
		CreateRole(m *models.RoleItems) (err error)
	}
)

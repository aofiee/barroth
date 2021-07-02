package domains

import (
	"github.com/aofiee/barroth/models"
)

type (
	ModuleUseCase interface {
		CreateModule(m *models.Modules) (err error)
		GetModule(m *models.Modules, slug string) (err error)
		UpdateModule(m *models.Modules, id string) (err error)
	}
	ModuleRepository interface {
		CreateModule(m *models.Modules) (err error)
		GetModule(m *models.Modules, slug string) (err error)
		UpdateModule(m *models.Modules, id string) (err error)
	}
)

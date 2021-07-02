package usecases

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
)

type (
	moduleUseCase struct {
		moduleRepo domains.ModuleRepository
	}
)

func NewModuleUseCase(repo domains.ModuleRepository) domains.ModuleUseCase {
	return &moduleUseCase{
		moduleRepo: repo,
	}
}
func (m *moduleUseCase) CreateModule(module *models.Modules) error {
	err := m.moduleRepo.CreateModule(module)
	return err
}
func (m *moduleUseCase) UpdateModule(module *models.Modules, id string) error {
	var chk models.Modules
	err := m.moduleRepo.GetModule(&chk, id)
	if err != nil {
		return err
	}
	err = m.moduleRepo.UpdateModule(module, id)
	return err
}
func (m *moduleUseCase) GetModule(module *models.Modules, slug string) error {
	err := m.moduleRepo.GetModule(module, slug)
	return err
}

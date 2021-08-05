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
	if err != nil {
		return err
	}
	roles, err := m.moduleRepo.GetAllRoles()
	if err != nil {
		return err
	}
	for _, v := range roles {
		exec := 0
		if v.Name == "Administrator" {
			exec = 1
		}
		m.moduleRepo.SetPermission(module.ID, v.ID, exec)
	}
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
func (m *moduleUseCase) GetModuleBySlug(module *models.Modules, method, slug string) error {
	err := m.moduleRepo.GetModuleBySlug(module, method, slug)
	return err
}
func (m *moduleUseCase) GetAllRoles() ([]models.RoleItems, error) {
	roles, err := m.moduleRepo.GetAllRoles()
	return roles, err
}
func (m *moduleUseCase) SetPermission(moduleID, roleID uint, exec int) error {
	err := m.moduleRepo.SetPermission(moduleID, roleID, exec)
	return err
}

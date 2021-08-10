package usecases

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
)

type (
	systemUseCase struct {
		systemRepo domains.SystemRepository
	}
)

func NewSystemUseCase(repo domains.SystemRepository) domains.SystemUseCase {
	return &systemUseCase{
		systemRepo: repo,
	}
}
func (s *systemUseCase) CreateSystem(system *models.System) error {
	err := s.systemRepo.CreateSystem(system)
	return err
}
func (s *systemUseCase) UpdateSystem(system *models.System, id string) error {
	var chk models.System
	err := s.systemRepo.GetSystem(&chk, id)
	if err != nil {
		return err
	}
	err = s.systemRepo.UpdateSystem(system, id)
	return err
}
func (s *systemUseCase) GetSystem(system *models.System, id string) error {
	err := s.systemRepo.GetSystem(system, id)
	return err
}
func (s *systemUseCase) GetFirstSystemInstallation(system *models.System) error {
	err := s.systemRepo.GetFirstSystemInstallation(system)
	return err
}
func (s *systemUseCase) CreateUser(m *models.Users) error {
	err := s.systemRepo.HashPassword(m)
	if err != nil {
		return err
	}
	err = s.systemRepo.CreateUser(m)
	return err
}
func (s *systemUseCase) CreateRole(m *models.RoleItems) error {
	err := s.systemRepo.CreateRole(m)
	return err
}
func (s *systemUseCase) SetExecToAllModules(m *[]models.Modules, roleID uint, isExec int) error {
	err := s.systemRepo.GetAllModules(m)
	if err != nil {
		return err
	}
	for _, v := range *m {
		var permissions models.Permissions
		permissions.ModuleID = v.ID
		permissions.RoleItemID = roleID
		permissions.IsExec = &isExec
		err := s.systemRepo.SetPermissions(&permissions)
		if err != nil {
			return err
		}
	}
	return nil
}

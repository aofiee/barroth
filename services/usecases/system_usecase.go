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

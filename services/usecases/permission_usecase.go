package usecases

import (
	"github.com/aofiee/barroth/domains"
	"github.com/aofiee/barroth/models"
)

type (
	permissionsUseCase struct {
		permissionsRepo domains.PermissionsRepository
	}
)

func NewPermissionsUseCase(repo domains.PermissionsRepository) domains.PermissionsUseCase {
	return &permissionsUseCase{
		permissionsRepo: repo,
	}
}
func (p *permissionsUseCase) SetPermissions(m *[]models.Permissions) error {
	for _, v := range *m {
		err := p.permissionsRepo.SetPermissions(&v)
		if err != nil {
			return err
		}
	}
	return nil
}

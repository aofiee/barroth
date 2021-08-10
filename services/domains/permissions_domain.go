package domains

import "github.com/aofiee/barroth/models"

type (
	PermissionsUseCase interface {
		SetPermissions(m *[]models.Permissions) (err error)
	}
	PermissionsRepository interface {
		SetPermissions(m *models.Permissions) (err error)
	}
)

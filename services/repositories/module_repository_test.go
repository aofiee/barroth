package repositories

import (
	"reflect"
	"testing"

	"github.com/aofiee/barroth/databases"
	"github.com/aofiee/barroth/models"
	"github.com/stretchr/testify/assert"
)

func TestModuleRepo(t *testing.T) {
	SetupMock(t)
	t.Run("TEST_MODULE_REPO", func(t *testing.T) {
		repo := NewModuleRepository(databases.DB)
		assert.Equal(t, "*repositories.moduleRepository", reflect.TypeOf(repo).String(), "new repo")
		m := models.Modules{
			Name:        "test",
			Description: "test",
			ModuleSlug:  "test",
		}
		err := repo.CreateModule(&m)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		err = repo.UpdateModule(&m, "1")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
		err = repo.GetModule(&m, "test")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
	})
}

package routes

import (
	"reflect"
	"testing"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/stretchr/testify/assert"
)

func TestInitAllRoutes(t *testing.T) {
	SetupMock(t)
	app := InitAllRoutes()
	assert.Equal(t, "*fiber.App", reflect.TypeOf(app).String(), "TestInitAllRoutes")
	r := NewRoleRoutes(barroth_config.ENV)
	r.Install(app)
}

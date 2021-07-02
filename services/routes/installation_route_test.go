package routes

import (
	"reflect"
	"testing"

	barroth_config "github.com/aofiee/barroth/config"
	"github.com/stretchr/testify/assert"
)

func TestInstallation(t *testing.T) {
	ins := NewInstallationRoutes(barroth_config.ENV)
	assert.Equal(t, "*routes.installationRoutes", reflect.TypeOf(ins).String(), "new installation")
	app := ins.Setup()
	ins.Install(app)
}

package config

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	t.Run("TEST_FAIL_CONFIG", func(t *testing.T) {
		config, err := LoadConfig("./")
		if err != nil {
			assert.NotEqual(t, nil, err, "config")
		}
		assert.Equal(t, "config.Config", reflect.TypeOf(config).String(), "config")
		assert.NotEqual(t, nil, config, "config")
	})
	t.Run("TEST_SUCCESS_CONFIG", func(t *testing.T) {
		config, err := LoadConfig("../")
		if err != nil {
			assert.NotEqual(t, nil, err, "config")
		}
		assert.Equal(t, "config.Config", reflect.TypeOf(config).String(), "config")
		assert.NotEqual(t, nil, config, "config")
	})
}

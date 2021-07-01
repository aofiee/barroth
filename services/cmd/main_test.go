package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetupDatabase(t *testing.T) {
	t.Run("TEST_LOAD_FAIL_ENV", func(t *testing.T) {
		dns, err := setupDNSDatabaseConnection("./")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
			assert.Equal(t, "", dns, "dns connection is empty")
		}
	})
	t.Run("TEST_LOAD_SUCCESS_ENV", func(t *testing.T) {
		dns, err := setupDNSDatabaseConnection("../")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
			assert.Equal(t, "", dns, "dns connection is empty")
		}
		assert.NotEqual(t, "", dns, "dns connection is empty")
	})
	t.Run("DATABASE_TEST", func(t *testing.T) {
		dns, err := setupDNSDatabaseConnection("../")
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
			assert.Equal(t, "", dns, "dns connection is empty")
		}
		err = createDatabaseConnection(dns)
		if err != nil {
			assert.NotEqual(t, nil, err, err.Error())
		}
	})
}

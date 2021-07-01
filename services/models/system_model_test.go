package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	s := System{}
	tb := s.TableName()
	assert.Equal(t, "system", tb, "TestTableName")
}

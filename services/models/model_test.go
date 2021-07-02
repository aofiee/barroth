package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTableName(t *testing.T) {
	s := System{}
	tb := s.TableName()
	assert.Equal(t, "system", tb, "TestTableName")

	users := Users{}
	tb = users.TableName()
	assert.Equal(t, "users", tb, "TestTableName")

	userroles := UserRoles{}
	tb = userroles.TableName()
	assert.Equal(t, "user_roles", tb, "TestTableName")

	permissions := Permissions{}
	tb = permissions.TableName()
	assert.Equal(t, "permissions", tb, "TestTableName")

	module := Modules{}
	tb = module.TableName()
	assert.Equal(t, "modules", tb, "TestTableName")

	roleitems := RoleItems{}
	tb = roleitems.TableName()
	assert.Equal(t, "role_items", tb, "TestTableName")
}

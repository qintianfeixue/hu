// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"ceshi1/app/dao/internal"
)

// roleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type roleDao struct {
	internal.RoleDao
}

var (
	// Role is globally public accessible object for table role operations.
	Role = roleDao{
		internal.Role,
	}
)

// Fill with you ideas below.

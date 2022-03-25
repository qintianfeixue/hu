package dao

import (
	"ceshi1/app/dao/internal"
)

// omsOrderDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type omsOrderDao struct {
	internal.OmsOrderDao
}

var (
	// OmsOrder is globally public accessible object for table role_permission operations.
	OmsOrder = omsOrderDao{
		internal.OmsOrder,
	}
)

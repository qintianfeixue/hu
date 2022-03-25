package model

import (
	"ceshi1/app/model/internal"
)

// OrderUser is the golang structure for table oms_order.
type OrderUser internal.OrderUser

// Fill with you ideas below.

type OrderUsers OrderUser

//连表查询oms_order 和 users 表
type OrderAndUser struct {
	Id             int
	MemberId       int
	CouponId       int
	status         int
	OrderSn        string
	MemberUsername string
	TotalAmount    string
	PayAmount      string
	FreightAmount  string
	NickName       string
	Mobile         string
}

type OrderAndUserLimitAll struct {
	OmsOrder       []OrderUser
	Page           int `json:"page" `
	Num            int `json:"num" `
	Count          int
	Next           bool
	Status         int    `json:"status" `
	MemberId       int    `json:"memberId" `
	MemberUsername string `json:"memberUsername" `
	OrderSn        string `json:"orderSn" `
}

type OrderAndUserLimit struct {
	OrderUser []OrderUser
	Page      int `json:"page" `
	Num       int `json:"num" `
	Count     int
	Next      bool
	//Status	       int `json:"status" `
	//MemberId       int `json:"memberId" `
	//MemberUsername string `json:"memberUsername" `
	//OrderSn        string `json:"orderSn" `
}

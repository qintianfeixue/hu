package model

import (
	"ceshi1/app/model/internal"
)

// OmsOrder is the golang structure for table oms_order.
type OmsOrder internal.OmsOrder

// Fill with you ideas below.

type OmsOrders OmsOrder

type OrderCha struct {
	Id             int
	MemberId       int
	CouponId       int
	status         int
	OrderSn        string
	MemberUsername string
	TotalAmount    string
	PayAmount      string
	FreightAmount  string
}

type OrderLimitCha struct {
	OmsOrder       []OmsOrder
	Page           int `json:"page" `
	Num            int `json:"num" `
	Count          int
	Next           bool
	Status         int    `json:"status" `
	MemberId       int    `json:"memberId" `
	MemberUsername string `json:"memberUsername" `
	OrderSn        string `json:"orderSn" `
}

type OrderLimitCha1 struct {
	OmsOrder []OmsOrder
	Page     int `json:"page" `
	Num      int `json:"num" `
	Count    int
	Next     bool
	//Status	       int `json:"status" `
	//MemberId       int `json:"memberId" `
	//MemberUsername string `json:"memberUsername" `
	//OrderSn        string `json:"orderSn" `
}

package internal

type OrderUser struct {
	Id             int    `orm:"id"       json:"id"`
	MemberId       int    `orm:"member_id"        json:"memberId"`
	CouponId       int    `orm:"coupon_id"        json:"couponId"`
	OrderSn        string `orm:"order_sn"      json:"orderSn"`
	MemberUsername string `orm:"member_username"    json:"memberUsername"`
	TotalAmount    string `orm:"total_amount"  json:"totalAmount"`
	PayAmount      string `orm:"pay_amount"    json:"payAmount"`
	FreightAmount  string `orm:"freight_amount"     json:"freightAmount"`
	Status         int    `orm:"status"     json:"status"`
	NickName       string `orm:"nickname"     json:"nickName"`
	Mobile         string `orm:"mobile"       json:"mobile"`
}

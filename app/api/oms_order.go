package api

import (
	"ceshi1/app/model"
	"ceshi1/app/service"
	"ceshi1/library/response"
	_ "ceshi1/library/response"
	"fmt"
	_ "github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/net/ghttp"
)

// 用户认证API管理对象
var AuthOmsOrder = new(authOmsOrderApi)

type authOmsOrderApi struct{}

func (a *authOmsOrderApi) OrderAll1(r *ghttp.Request) {

	var res []*model.OmsOrder

	res, _ = service.OmsOrder.Chao()
	if res != nil {
		r.Response.WriteJson(res)
		r.ExitAll()
	}

}

func (a *authOmsOrderApi) OrderAll(r *ghttp.Request) {
	var (
		data *model.OrderLimitCha
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var res []*model.OmsOrder

	res, _ = service.OmsOrder.Chao()
	//service.OmsOrder.Limit123o(r.Context(), data.Page, data.Num)
	var sum int
	sum = len(res)
	fmt.Println(sum)
	var i int
	var res1 []*model.OmsOrder
	res1, _ = service.OmsOrder.Limit123o(r.Context(), data.Page, data.Num)
	var u model.OrderLimitCha
	u.Count = sum
	//
	u.OmsOrder = make([]model.OmsOrder, len(res1))
	u.Page = data.Page
	u.Num = data.Num

	for i = 0; i < len(res1); i++ {

		//fmt.Println(user)
		u.OmsOrder[i] = *res1[i]
	}
	//fmt.Println((u.Count-(u.Page*u.Num))/u.Num)
	//要把本页的值也加上
	if u.Count-(u.Page*u.Num+u.Num) > 0 {
		u.Next = true
	} else {
		u.Next = false
	}

	r.Response.WriteJson(u)

}

func (a *authOmsOrderApi) OrderAll2(r *ghttp.Request) {
	var (
		data *model.OrderLimitCha
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var res []model.OmsOrder

	res, _ = service.OmsOrder.Cha1(r.Context(), data)
	//service.OmsOrder.Limit123o(r.Context(), data.Page, data.Num)
	var sum int
	sum = len(res)

	fmt.Println(sum)

	var i int
	var res1 []model.OmsOrder
	res1, _ = service.OmsOrder.Limit123od(r.Context(), data)
	//fmt.Println(res1)
	var u model.OrderLimitCha1

	//
	u.OmsOrder = make([]model.OmsOrder, len(res1))
	u.Page = data.Page
	u.Num = data.Num
	u.Count = sum
	for i = 0; i < len(res1); i++ {
		//fmt.Println(user)
		u.OmsOrder[i] = res1[i]
	}
	//fmt.Println((u.Count-(u.Page*u.Num))/u.Num)
	//要把本页的值也加上
	if u.Count-(u.Page*u.Num+u.Num) > 0 {
		u.Next = true
	} else {
		u.Next = false
	}
	r.Response.WriteJson(u)

}

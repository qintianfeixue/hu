package api

import (
	"ceshi1/app/model"
	"ceshi1/app/service"
	"ceshi1/library/response"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
)

// 用户认证API管理对象
var OrderUser = new(orderUserApi)

type orderUserApi struct{}

func (a *orderUserApi) OrderUserAll(r *ghttp.Request) {

	var res []model.OrderUser

	res, _ = service.OrderUser.ChaOrderUser()
	if res != nil {
		r.Response.WriteJson(res)
		r.ExitAll()
	}

}

func (a *orderUserApi) OrderUserLimit(r *ghttp.Request) {
	var (
		data *model.OrderAndUserLimitAll
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var res []model.OrderUser

	res, _ = service.OrderUser.ChaOrderUserAll(r.Context(), data)

	var sum int
	sum = len(res)

	fmt.Println(sum)

	var i int
	var res1 []model.OrderUser
	res1, _ = service.OrderUser.LimitOrderUser(r.Context(), data)
	//fmt.Println(res1)
	var u model.OrderAndUserLimit

	//
	u.OrderUser = make([]model.OrderUser, len(res1))
	u.Page = data.Page
	u.Num = data.Num
	u.Count = sum
	for i = 0; i < len(res1); i++ {
		//fmt.Println(user)
		u.OrderUser[i] = res1[i]
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

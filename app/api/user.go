package api

import (
	"ceshi1/app/service"
	"ceshi1/library/response"
	"fmt"

	"github.com/gogf/gf/net/ghttp"
)

// 用户api管理对象
var User = new(userApi)

type userApi struct{}

// @summary 获取用户信息
// @tags    用户服务
// @produce json
// @param   token formData string true "token"
// @router  /oa/login [GET]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *userApi) UserInfo(r *ghttp.Request) {
	fmt.Println("____________")
	user_info, err := service.User.UserInfo(r.Context())
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", user_info)
}

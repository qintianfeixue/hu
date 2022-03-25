package api

import (
	"ceshi1/app/model"
	"ceshi1/app/service"
	"ceshi1/library/response"
	"fmt"

	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/net/ghttp"
)

// 用户认证API管理对象
var AuthUser = new(authUserApi)

type authUserApi struct{}

// @summary 用户登录验证
// @tags    用户token
// @produce json
// @param   userName formData string true "用户账号"
// @param   userPassword formData string true "用户密码"
// @router  /oa/login [POST]
// @success 200 {object} response.JsonResponse "执行结果"
//func (a *authUserApi) LoginBefore(r *ghttp.Request) (string, interface{}) {
//	var (
//		data *model.UserApiSignInReq
//	)
//	if err := r.Parse(&data); err != nil {
//		response.JsonExit(r, 1, err.Error())
//	}
//	if err := service.User.SignIn(r.Context(), data.UserName, data.UserPassword); err != nil {
//		r.Response.WriteJson(gtoken.Fail(err.Error()))
//		r.ExitAll()
//	}
//	return data.UserName, nil
//}

//查询所有
func (a *authUserApi) LoginBefore1(r *ghttp.Request) (string, interface{}) {
	var (
		data *model.UserApiSignInReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	var res []*model.User

	res, _ = service.User.Cha()
	if res != nil {
		r.Response.WriteJson(res)
		r.ExitAll()
	}

	return data.UserName, nil

}
func (a *authUserApi) LoginBefore(r *ghttp.Request) (string, interface{}) {
	var (
		data *model.UserAuth
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var res []*model.User

	res, _ = service.User.Cha()

	var sum int
	sum = len(res)

	var i int
	var res1 []*model.User
	res1, _ = service.User.Limit123(r.Context(), data.Page, data.Num)
	var u model.UserAuth
	u.Count = sum
	//
	u.User = make([]model.User, len(res1))
	u.Page = data.Page
	u.Num = data.Num

	for i = 0; i < len(res1); i++ {
		user := res[i]
		//fmt.Println(user)
		u.User[i] = *user
	}
	//fmt.Println((u.Count-(u.Page*u.Num))/u.Num)
	if u.Count-(u.Page*u.Num) > 0 {
		u.Next = (u.Count - (u.Page * u.Num)) / u.Num
		if (u.Count-(u.Page*u.Num))%u.Num > 0 {
			u.Next = u.Next + 1
		}
	}

	r.Response.WriteJson(u)
	return u.User[0].UserName, nil

}

func (a *authUserApi) Test(r *ghttp.Request) {
	var (
		data *model.UserAuth
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}

	var res []*model.User

	res, _ = service.User.Cha()

	var sum int
	sum = len(res)

	var i int
	var res1 []*model.User
	res1, _ = service.User.Limit123(r.Context(), data.Page, data.Num)
	var u model.UserAuth
	u.Count = sum

	//切片实现动态数组
	u.User = make([]model.User, len(res1))
	u.Page = data.Page
	u.Num = data.Num

	for i = 0; i < len(res1); i++ {
		user := res[i]
		//fmt.Println(user)
		u.User[i] = *user
	}
	//fmt.Println((u.Count-(u.Page*u.Num))/u.Num)
	if u.Count-(u.Page*u.Num) > 0 {
		u.Next = (u.Count - (u.Page * u.Num)) / u.Num
		if (u.Count-(u.Page*u.Num))%u.Num > 0 {
			u.Next = u.Next + 1
		}
	}

	r.Response.WriteJson(u)

}

//分页查询
func (a *authUserApi) LoginBefore2(r *ghttp.Request) (string, interface{}) {
	var (
		data *model.UserApiLimit
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	var res []*model.User

	res, _ = service.User.Limit12(r.Context(), data.Page, data.Num)
	if res != nil {
		r.Response.WriteJson(res)
		r.ExitAll()
	}
	return data.Page, nil

}

// @summary 登录返回方法
// @tags    用户token
// @produce json
// @param   userName formData string true "用户账号"
// @param   userPassword formData string true "用户密码"
// @router  /oa/login [POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *authUserApi) LoginAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		//格式原因暂时注释掉

		//token := respData.GetString("token")
		//// user_info, _ := service.User.UserInfo(r.Context())
		//response.JsonExit(r, 0, "用户登录成功", g.Map{
		//	"token": token,
		//	// "userInfo": user_info,
		//})
	}
}

// @summary 后台退出登录
// @tags    用户token
// @produce json
// @param   token formData string true "token"
// @router  /oa/logot [GET , POST]
// @success 200 {object} response.JsonResponse "执行结果"
func (a *authUserApi) LogoutAfter(r *ghttp.Request, respData gtoken.Resp) {
	if !respData.Success() {
		r.Response.WriteJson(respData)
	} else {
		response.JsonExit(r, 0, "用户退出成功")
	}
}

// @summary token认证成功后调用
// @tags    用户token
// @produce json
// @param   token formData string true "token"
// @success 200 {object} response.JsonResponse "执行结果"
func (a *authUserApi) AuthAfter(r *ghttp.Request, respData gtoken.Resp) {
	fmt.Println("token认证成功后调用")
	if r.Method == "OPTIONS" || respData.Success() {
		userName := respData.GetString("userKey")
		service.User.GetUserName(r.Context(), userName)
		r.Middleware.Next()
	} else {
		respData.Msg = "用户信息验证失败"
		response := r.Response
		options := response.DefaultCORSOptions()
		response.CORS(options)
		response.WriteJson(respData)
		r.ExitAll()
	}
}

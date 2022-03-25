package router

import (
	"ceshi1/app/api"
	"ceshi1/app/service"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

//路由改变之后服务要重启
func init() {
	urlPath := g.Config().GetString("url_path")
	s := g.Server()
	//直接在根下加  /test
	s.BindHandler("/test", api.AuthUser.Test)
	s.BindHandler("/", func(r *ghttp.Request) {

	})
	// 前台系统自定义错误页面
	s.BindStatusHandler(403, func(r *ghttp.Request) {
		r.Response.RedirectTo("/login")
	})

	//分组路由注册方式
	s.Group(urlPath+"/", func(group *ghttp.RouterGroup) {
		//自定义上下文和允许跨域
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,
		)

		group.GET("/test", api.AuthUser.Test)
		group.GET("/order", api.AuthOmsOrder.OrderAll)
		group.GET("/order1", api.AuthOmsOrder.OrderAll1)
		group.GET("/order2", api.AuthOmsOrder.OrderAll2)
		group.GET("/orderU", api.OrderUser.OrderUserAll)
		group.GET("/orderL", api.OrderUser.OrderUserLimit)
		gfToken := &gtoken.GfToken{
			CacheMode:       1,                        //1 gcache 2 gredis 默认1
			Timeout:         60 * 60 * 1000,           //默认10天（毫秒）
			LoginPath:       "/login",                 //登录接口路径
			LoginBeforeFunc: api.AuthUser.LoginBefore, //登录验证需要用户实现方法
			LoginAfterFunc:  api.AuthUser.LoginAfter,
			LogoutPath:      "/logout", //登出接口路径
			// LogoutBeforeFunc: common.LogoutBefore,//登出接口前调用
			LogoutAfterFunc: api.AuthUser.LogoutAfter,
			AuthAfterFunc:   api.AuthUser.AuthAfter,
		}
		group.Group("/", func(group *ghttp.RouterGroup) {
			// group.Middleware(service.Middleware.Auth)

			gfToken.Middleware(group)
			group.ALL("/index", func(r *ghttp.Request) {
				r.Response.Writeln("11111111")
			})

			// group.ALL("/getUserInfo", api.User.UserInfo)
			group.ALL("/getUserInfo", api.Index.GetInfo)
			group.ALL("/getRouters", api.Index.GetRouters)
			// group.Group("/user", func(group *ghttp.RouterGroup) {
			// 	group.POST("/getUserInfo", api.User.UserInfo)
			// })
		})
	})
}

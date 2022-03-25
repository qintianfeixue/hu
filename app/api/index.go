package api

import (
	"ceshi1/app/service"
	"ceshi1/library/response"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
)

// 公共api管理对象
var Index = new(indexApi)

type indexApi struct{}

// @summary 获取用户信息
// @tags    公共
// @produce json
// @param   token formData string true "token"
// @router  /oa/index/getInfo [GET]
// @success 0 {object} response.JsonResponse "执行结果"
func (a *indexApi) GetInfo(r *ghttp.Request) {
	//获取用户信息
	user_entity, err := service.User.UserInfo(r.Context())
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	user_info := gconv.Map(user_entity)
	if user_entity == nil {
		response.JsonExit(r, 1, "查询用户信息失败！")
	}
	// auth_rp, err := service.AuthRule.GetUserRole(user_entity.Id)
	// if err != nil {
	// 	response.JsonExit(r, 1, err.Error())
	// }

	// result := g.Map{
	// 	"user":        user_info,
	// 	"roles":       auth_rp["role"],
	// 	"permissions": auth_rp["permission"],
	// }
	// 获取用户角色信息
	role_entity, err := service.AuthRule.UserRoleInfoGet(user_entity.Id)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	role_id_list := make([]int, len(role_entity))
	role_name_list := make([]string, len(role_entity))
	for k, v := range role_entity {
		role_id_list[k] = int(v.Role.Id)
		role_name_list[k] = v.Role.Name
	}
	// 获取角色对应菜单信息
	per_entity, err := service.AuthRule.RolePerInfoGet(role_id_list)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	per_name_list := make([]string, len(per_entity))
	for k, v := range per_entity {
		per_name_list[k] = v.Permission.Name
	}

	result := g.Map{
		"user":        user_info,
		"roles":       role_name_list,
		"permissions": per_name_list,
	}
	response.JsonExit(r, 0, "操作成功", result)
}

// @summary 获取后台菜单
// @tags    公共
// @produce json
// @param   token formData string true "token"
// @router  /oa/index/getInfo [GET]
// @success 0 {object} response.JsonResponse "执行结果"
func (a *indexApi) GetRouters(r *ghttp.Request) {
	// 获取用户信息
	user_entity, err := service.User.UserInfo(r.Context())
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	// 判断是否是管理员，管理员无需验证权限
	if user_entity.IsAdmin == "1" {
		per_entity, err := service.AuthRule.PerTreeInfoAll()
		if err != nil {
			response.JsonExit(r, 1, err.Error())
		}
		response.JsonExit(r, 0, "ok", per_entity)
	}
	// 获取用户角色
	role_entity, err := service.AuthRule.UserRoleInfoGet(user_entity.Id)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	role_id_list := make([]int, len(role_entity)) // 角色id列表
	for k, v := range role_entity {
		role_id_list[k] = int(v.Role.Id)
	}
	// 获取角色菜单
	per_entity, err := service.AuthRule.RolePerTreeInfoGet(role_id_list)
	if err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	response.JsonExit(r, 0, "ok", per_entity)
}

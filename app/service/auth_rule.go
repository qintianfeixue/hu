package service

import (
	"ceshi1/app/dao"
	"ceshi1/app/model"
	"fmt"
)

// 身份验证规则
var AuthRule = new(authRuleService)

type authRuleService struct{}

// 获取用户对应的角色name
// func (a *authRuleService) GetUserRole(userId int) (map[string][]string, error) {
// 	// 获取对应的角色id
// 	user_role_entity, err := dao.UserRole.FindAll("user_id =?", userId)
// 	roles_list := make([]string, len(user_role_entity))
// 	rolei_list := make([]int, len(user_role_entity))
// 	for k, v := range user_role_entity {
// 		// 获取角色name
// 		role_entity, _ := dao.Role.FindOne("id=?", v.RoleId)
// 		roles_list[k] = role_entity.Name
// 		rolei_list[k] = int(role_entity.Id)
// 		// 获取菜单权限
// 		// pers_list, _ := a.GetRolePermission(int(role_entity.Id))
// 		// fmt.Println(pers_list)id_list:=
// 	}
// 	pern_list, _ := a.GetRolePermission(rolei_list)
// 	auth_rp := map[string][]string{
// 		"role":       roles_list,
// 		"permission": pern_list,
// 	}
// 	return auth_rp, err
// }

// // 获取角色对应菜单权限
// func (a *authRuleService) GetRolePermission(roleId []int) ([]string, error) {
// 	role_permission_entity, _ := dao.RolePermission.All("role_id", roleId)
// 	perid_list := make([]int, len(role_permission_entity))
// 	for k, v := range role_permission_entity {
// 		perid_list[k] = v.PermissionId
// 	}

// 	permission_entity, err := dao.Permission.FindAll("id", perid_list)
// 	pern_list := make([]string, len(permission_entity))
// 	for k2, v2 := range permission_entity {
// 		pern_list[k2] = v2.Name
// 	}
// 	return pern_list, err
// }

// 根据用户id，获取用户对应角色信息——新
func (a *authRuleService) UserRoleInfoGet(userId int) (role []*model.UserRoleGetOne, err error) {
	err = dao.UserRole.Where("user_id=?", userId).ScanList(&role, "UserRole")
	// 通过name查询
	// err = dao.UserRole.Where("user_name=?", userName).ScanList(&role, "UserRole")
	if err != nil {
		return
	}
	err = dao.Role.Where("status=1").ScanList(&role, "Role", "UserRole", "id:roleId")
	if err != nil {
		return
	}
	return
}

// 根据角色id，获取角色对应菜单信息——新
func (a *authRuleService) RolePerInfoGet(roleId []int) (role_per []*model.RolePerGetOne, err error) {
	err = dao.RolePermission.Where("role_id IN(?)", roleId).ScanList(&role_per, "RolePermission")
	if err != nil {
		return
	}
	err = dao.Permission.Where("status").ScanList(&role_per, "Permission", "RolePermission", "id:permissionId")
	return
}

// 根据角色id，获取菜单信息，树形结构
func (a *authRuleService) RolePerTreeInfoGet(roleId []int) (role_per []*model.PerServiceGetOne, err error) {
	var role_per_entity []*model.RolePermission
	role_per_entity, err = dao.RolePermission.FindAll("role_id IN(?)", roleId)
	if err != nil {
		return
	}
	per_id_list := make([]int, len(role_per_entity))
	for k, v := range role_per_entity {
		per_id_list[k] = v.PermissionId
	}
	fmt.Println(per_id_list)
	role_per, err = a.PerTreeInfoGet(per_id_list)
	return
}

// 根据菜单id，获取菜单详细信息,树形结构——新
func (a *authRuleService) PerTreeInfoGet(perId []int) (menus []*model.PerServiceGetOne, err error) {
	// 获取主目录
	err = dao.Permission.Where("id IN(?) AND pid=0", perId).ScanList(&menus, "Permission")
	if err != nil {
		return
	}
	// 获取子目录
	err = dao.Permission.Where("id IN(?) AND pid!=0", perId).ScanList(&menus, "Children", "Permission", "pid:id")
	if err != nil {
		return
	}
	return menus, err
}

// 管理员获取所有菜单信息，树形结构
func (a *authRuleService) PerTreeInfoAll() (menus []*model.PerServiceGetOne, err error) {
	// 获取主目录
	err = dao.Permission.Where("pid=0").ScanList(&menus, "Permission")
	if err != nil {
		return
	}
	// 获取子目录
	err = dao.Permission.Where("pid!=0").ScanList(&menus, "Children", "Permission", "pid:id")
	if err != nil {
		return
	}
	return menus, err
}

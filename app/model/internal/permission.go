// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

// Permission is the golang structure for table permission.
type Permission struct {
	Id         uint   `orm:"id,primary"  json:"id"`         //
	Pid        uint   `orm:"pid"         json:"pid"`        // 父ID
	Name       string `orm:"name,unique" json:"name"`       // 规则名称
	Title      string `orm:"title"       json:"title"`      // 菜单名称
	Icon       string `orm:"icon"        json:"icon"`       // 图标
	Condition  string `orm:"condition"   json:"condition"`  // 条件（没用）
	Remark     string `orm:"remark"      json:"remark"`     // 备注
	MenuType   uint   `orm:"menu_type"   json:"menuType"`   // 类型 0目录 1菜单 2按钮
	Createtime uint   `orm:"createtime"  json:"createtime"` // 创建时间
	Updatetime uint   `orm:"updatetime"  json:"updatetime"` // 更新时间
	Weigh      int    `orm:"weigh"       json:"weigh"`      // 权重 越低越靠前
	Status     uint   `orm:"status"      json:"status"`     // 状态 0停用 1正常
	AlwaysShow uint   `orm:"always_show" json:"alwaysShow"` // 显示状态（没用）0隐藏 1显示
	Path       string `orm:"path"        json:"path"`       // 路由地址
	Component  string `orm:"component"   json:"component"`  // 组件路径
	IsFrame    uint   `orm:"is_frame"    json:"isFrame"`    // 是否外链 1是 0否
	ModuleType string `orm:"module_type" json:"moduleType"` // 所属模块（没用）
	ModelId    uint   `orm:"model_id"    json:"modelId"`    // 模型ID（没用）
}
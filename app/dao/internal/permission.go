// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"time"

	"ceshi1/app/model"
)

// PermissionDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type PermissionDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns permissionColumns
}

// PermissionColumns defines and stores column names for table permission.
type permissionColumns struct {
	Id         string //
	Pid        string // 父ID
	Name       string // 规则名称
	Title      string // 菜单名称
	Icon       string // 图标
	Condition  string // 条件（没用）
	Remark     string // 备注
	MenuType   string // 类型 0目录 1菜单 2按钮
	Createtime string // 创建时间
	Updatetime string // 更新时间
	Weigh      string // 权重 越低越靠前
	Status     string // 状态 0停用 1正常
	AlwaysShow string // 显示状态（没用）0隐藏 1显示
	Path       string // 路由地址
	Component  string // 组件路径
	IsFrame    string // 是否外链 1是 0否
	ModuleType string // 所属模块（没用）
	ModelId    string // 模型ID（没用）
}

var (
	// Permission is globally public accessible object for table permission operations.
	Permission = PermissionDao{
		M:     g.DB("default").Model("permission").Safe(),
		DB:    g.DB("default"),
		Table: "permission",
		Columns: permissionColumns{
			Id:         "id",
			Pid:        "pid",
			Name:       "name",
			Title:      "title",
			Icon:       "icon",
			Condition:  "condition",
			Remark:     "remark",
			MenuType:   "menu_type",
			Createtime: "createtime",
			Updatetime: "updatetime",
			Weigh:      "weigh",
			Status:     "status",
			AlwaysShow: "always_show",
			Path:       "path",
			Component:  "component",
			IsFrame:    "is_frame",
			ModuleType: "module_type",
			ModelId:    "model_id",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *PermissionDao) Ctx(ctx context.Context) *PermissionDao {
	return &PermissionDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *PermissionDao) As(as string) *PermissionDao {
	return &PermissionDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *PermissionDao) TX(tx *gdb.TX) *PermissionDao {
	return &PermissionDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *PermissionDao) Master() *PermissionDao {
	return &PermissionDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *PermissionDao) Slave() *PermissionDao {
	return &PermissionDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *PermissionDao) Args(args ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *PermissionDao) LeftJoin(table ...string) *PermissionDao {
	return &PermissionDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *PermissionDao) RightJoin(table ...string) *PermissionDao {
	return &PermissionDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *PermissionDao) InnerJoin(table ...string) *PermissionDao {
	return &PermissionDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *PermissionDao) Fields(fieldNamesOrMapStruct ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *PermissionDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *PermissionDao) Option(option int) *PermissionDao {
	return &PermissionDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *PermissionDao) OmitEmpty() *PermissionDao {
	return &PermissionDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *PermissionDao) Filter() *PermissionDao {
	return &PermissionDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *PermissionDao) Where(where interface{}, args ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *PermissionDao) WherePri(where interface{}, args ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *PermissionDao) And(where interface{}, args ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *PermissionDao) Or(where interface{}, args ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *PermissionDao) Group(groupBy string) *PermissionDao {
	return &PermissionDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *PermissionDao) Order(orderBy ...string) *PermissionDao {
	return &PermissionDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *PermissionDao) Limit(limit ...int) *PermissionDao {
	return &PermissionDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *PermissionDao) Offset(offset int) *PermissionDao {
	return &PermissionDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *PermissionDao) Page(page, limit int) *PermissionDao {
	return &PermissionDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *PermissionDao) Batch(batch int) *PermissionDao {
	return &PermissionDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *PermissionDao) Cache(duration time.Duration, name ...string) *PermissionDao {
	return &PermissionDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *PermissionDao) Data(data ...interface{}) *PermissionDao {
	return &PermissionDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*model.Permission.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *PermissionDao) All(where ...interface{}) ([]*model.Permission, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Permission
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *model.Permission.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *PermissionDao) One(where ...interface{}) (*model.Permission, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Permission
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *PermissionDao) FindOne(where ...interface{}) (*model.Permission, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *model.Permission
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *PermissionDao) FindAll(where ...interface{}) ([]*model.Permission, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*model.Permission
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *PermissionDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *PermissionDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *PermissionDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *PermissionDao) Chunk(limit int, callback func(entities []*model.Permission, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*model.Permission
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *PermissionDao) LockUpdate() *PermissionDao {
	return &PermissionDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *PermissionDao) LockShared() *PermissionDao {
	return &PermissionDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *PermissionDao) Unscoped() *PermissionDao {
	return &PermissionDao{M: d.M.Unscoped()}
}

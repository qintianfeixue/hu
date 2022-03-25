package service

import (
	_ "ceshi1/app/dao"
	"ceshi1/app/model"
	"context"
	"fmt"
	_ "fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	_ "strconv"
)

var OrderUser = new(OrderUserService)

type OrderUserService struct{}

// 订单信息及用户信息查询
func (s *OrderUserService) ChaOrderUser() (res []model.OrderUser, err error) {
	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("数据库连接失败")
	}
	orderAndUser := db.Model("oms_order o").LeftJoin("users u", "o.member_id=u.id").Fields("o.*,u.mobile,u.nickname")
	orderAndUser.FindAll()

	var result []model.OrderUser
	orderAndUser.Structs(&result)

	return result, nil
}

func (s *OrderUserService) LimitOrderUser(ctx context.Context, oms *model.OrderAndUserLimitAll) ([]model.OrderUser, error, int) {

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("数据库连接失败"), 0
	}
	orderAndUser := db.Model("oms_order o").LeftJoin("users u", "o.member_id=u.id").Fields("o.*,u.mobile,u.nickname")
	fmt.Println("fsfsfsaf:", oms.Status)
	if oms != nil {
		fmt.Println("我进来了")
		if oms.Status != 0 {
			fmt.Println("我进来了")
			orderAndUser.Where("status = ?", oms.Status)
		}
		if oms.OrderSn != "" {
			orderAndUser.Where("order_sn = ?", oms.OrderSn)
		}
		if oms.MemberId != 0 {
			orderAndUser.Where("member_id = ?", oms.MemberId)
		}
		if oms.MemberUsername != "" {
			orderAndUser.Where("member_username = ?", oms.MemberUsername)
		}
	}
	var result []model.OrderUser

	var count int
	count, _ = orderAndUser.Limit((oms.Page-1)*oms.Num, oms.Num).Count()

	orderAndUser.Limit((oms.Page-1)*oms.Num, oms.Num).FindAll()
	orderAndUser.Structs(&result)
	//res, err = dao.OmsOrder.Limit(start*end, end).FindAll()
	//res, err = dao.User.Limitya(start1,end1)
	return result, err, count
}

func (s *OrderUserService) ChaOrderUserAll(ctx context.Context, oms *model.OrderAndUserLimitAll) ([]model.OrderUser, error) {

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("数据库连接失败")
	}
	orderAndUser := db.Model("oms_order o").LeftJoin("users u", "o.member_id=u.id").Fields("o.*,u.mobile,u.nickname")
	fmt.Println("fsfsfsaf:", oms.Status)
	if oms != nil {
		fmt.Println("我进来了")
		if oms.Status != 0 {
			fmt.Println("我进来了")
			orderAndUser.Where("status = ?", oms.Status)
		}
		if oms.OrderSn != "" {
			orderAndUser.Where("order_sn = ?", oms.OrderSn)
		}
		if oms.MemberId != 0 {
			orderAndUser.Where("member_id = ?", oms.MemberId)
		}
		if oms.MemberUsername != "" {
			orderAndUser.Where("member_username = ?", oms.MemberUsername)
		}
	}
	var result []model.OrderUser

	orderAndUser.Structs(&result)
	//res, err = dao.OmsOrder.Limit(start*end, end).FindAll()
	//res, err = dao.User.Limitya(start1,end1)
	return result, err
}

package service

import (
	"ceshi1/app/dao"
	"ceshi1/app/model"
	"context"
	"fmt"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/errors/gerror"
	"strconv"
)

var OmsOrder = new(omsOderService)

type omsOderService struct{}

//分页查询实现
func (s *omsOderService) Limit12o(ctx context.Context, start, end string) (res []*model.OmsOrder, err error) {
	var start1 int
	start1, _ = strconv.Atoi(start)
	var end1 int
	end1, _ = strconv.Atoi(end)
	res, err = dao.OmsOrder.FindAll("id > 0 limit ?, ?", start1*end1, end1)
	//res, err = dao.User.Limitya(start1,end1)
	return res, err
}
func (s *omsOderService) Limit123o(ctx context.Context, start, end int) (res []*model.OmsOrder, err error) {

	res, err = dao.OmsOrder.Limit(start*end, end).FindAll()
	//res, err = dao.User.Limitya(start1,end1)
	return res, err
}

func (s *omsOderService) Limit123od(ctx context.Context, oms *model.OrderLimitCha) ([]model.OmsOrder, error) {

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("数据库连接失败")
	}
	omsOrder := db.Table("oms_order")
	fmt.Println("fsfsfsaf:", oms.Status)
	if oms != nil {
		fmt.Println("我进来了")
		if oms.Status != 0 {
			fmt.Println("我进来了")
			omsOrder.Where("status = ?", oms.Status)
		}
		if oms.OrderSn != "" {
			omsOrder.Where("order_sn = ?", oms.OrderSn)
		}
		if oms.MemberId != 0 {
			omsOrder.Where("member_id = ?", oms.MemberId)
		}
		if oms.MemberUsername != "" {
			omsOrder.Where("member_username = ?", oms.MemberUsername)
		}
	}
	var result []model.OmsOrder
	omsOrder.Limit((oms.Page-1)*oms.Num, oms.Num).FindAll()
	omsOrder.Structs(&result)
	//res, err = dao.OmsOrder.Limit(start*end, end).FindAll()
	//res, err = dao.User.Limitya(start1,end1)
	return result, err
}

func (s *omsOderService) Cha1(ctx context.Context, oms *model.OrderLimitCha) ([]model.OmsOrder, error) {

	db, err := gdb.Instance()
	if err != nil {
		return nil, gerror.New("数据库连接失败")
	}
	omsOrder := db.Table("oms_order")
	fmt.Println("fsfsfsaf:", oms.Status)
	if oms != nil {
		fmt.Println("我进来了")
		if oms.Status != 0 {
			fmt.Println("我进来了")
			omsOrder.Where("status = ?", oms.Status)
		}
		if oms.OrderSn != "" {
			omsOrder.Where("order_sn = ?", oms.OrderSn)
		}
		if oms.MemberId != 0 {
			omsOrder.Where("member_id = ?", oms.MemberId)
		}
		if oms.MemberUsername != "" {
			omsOrder.Where("member_username = ?", oms.MemberUsername)
		}
	}
	var result []model.OmsOrder

	omsOrder.Structs(&result)
	//res, err = dao.OmsOrder.Limit(start*end, end).FindAll()
	//res, err = dao.User.Limitya(start1,end1)
	return result, err
}

// 用户所有信息查询
func (s *omsOderService) Chao() (res []*model.OmsOrder, err error) {

	res, err = dao.OmsOrder.All()

	//dao.User.Limit(1).Scan(&res)
	return res, err
}

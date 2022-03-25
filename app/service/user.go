package service

import (
	"ceshi1/app/dao"
	"ceshi1/app/model"
	"context"
	"errors"
	"fmt"
	"strconv"
)

// 中间件管理服务
var User = new(userService)

type userService struct{}

// 判断用户是否已经登录,未使用
func (s *userService) IsSignedIn(ctx context.Context) bool {
	if v := Context.Get(ctx); v != nil && v.User != nil {
		return true
	}
	return false
}

// 用户登录，成功返回用户信息，否则返回nil; passport应当会md5值字符串
func (s *userService) SignIn(ctx context.Context, user_name, user_password string) error {
	user, err := dao.User.FindOne("user_name=? and user_password=?", user_name, user_password)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("账号或密码错误")
	}
	// if err := Session.SetUser(ctx, user); err != nil {
	// 	return nil
	// }
	Context.SetUser(ctx, &model.ContextUser{
		Id:       uint(user.Id),
		UserName: user.UserName,
	})
	return nil
}

//分页查询实现
func (s *userService) Limit12(ctx context.Context, start, end string) (res []*model.User, err error) {
	var start1 int
	start1, _ = strconv.Atoi(start)
	var end1 int
	end1, _ = strconv.Atoi(end)
	res, err = dao.User.FindAll("id > 0 limit ?, ?", start1*end1, end1)
	//res, err = dao.User.Limitya(start1,end1)
	return res, err
}
func (s *userService) Limit123(ctx context.Context, start, end int) (res []*model.User, err error) {

	res, err = dao.User.FindAll("id > 0 limit ?, ?", start*end, end)
	//res, err = dao.User.Limitya(start1,end1)
	return res, err
}

// 用户所有信息查询
func (s *userService) Cha() (res []*model.User, err error) {

	res, err = dao.User.All()

	//dao.User.Limit(1).Scan(&res)
	return res, err
}

// 用户token认证成功后，将username写入上下文
func (s *userService) GetUserName(ctx context.Context, userName string) {
	// customCtx := Context.Get(ctx)
	Context.SetUser(ctx, &model.ContextUser{
		UserName: userName,
	})

}

// 用户登录后，获取用户信息
func (s *userService) UserInfo(ctx context.Context) (res *model.UserApiGetInfo, err error) {
	customCtx := Context.Get(ctx)
	fmt.Println(customCtx.User.UserName)
	// m := dao.User.Fields(&model.UserApiGetInfo{})
	// return m.FindOne("id=?", customCtx.User.Id)
	// res = &model.UserApiGetInfo{}

	err = dao.User.Where("user_name=?", customCtx.User.UserName).Scan(&res)
	return res, err

	// return dao.User.FindOne("id=?", customCtx.User.Id)
}

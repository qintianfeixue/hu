package service

import (
	"ceshi1/app/model"
	"context"
)

//Session管理服务
var Session = new(sessionService)

type sessionService struct{}

const (
	//用户信息存放在Session中的key
	sessionKeyUser = "SessionKeyUser"
)

//设置用Session
func (s *sessionService) SetUser(ctx context.Context, user *model.User) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser, user)
}

//获取当前登录的用户信息对象，如果用户未登录返回nil
func (s *sessionService) GetUser(ctx context.Context) *model.User {
	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var user *model.User
			_ = v.Struct(&user)
			return user
		}
	}
	return nil
}

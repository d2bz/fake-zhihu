// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: user.proto

package server

import (
	"context"

	"zhihu/app/user_rpc/internal/logic"
	"zhihu/app/user_rpc/internal/svc"
	"zhihu/app/user_rpc/user"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) FindByMobile(ctx context.Context, in *user.FindByMobileRequest) (*user.FindByMobileResponse, error) {
	l := logic.NewFindByMobileLogic(ctx, s.svcCtx)
	return l.FindByMobile(in)
}

func (s *UserServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *UserServer) FindById(ctx context.Context, in *user.FindByIdRequest) (*user.FindByIdResponse, error) {
	l := logic.NewFindByIdLogic(ctx, s.svcCtx)
	return l.FindById(in)
}

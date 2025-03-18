package logic

import (
	"context"

	"zhihu/app/user_rpc/internal/svc"
	"zhihu/app/user_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByMobileLogic) FindByMobile(in *user.FindByMobileRequest) (*user.FindByMobileResponse, error) {
	u, err := l.svcCtx.UserModel.FindByMobile(l.ctx, in.Mobile)
	if err != nil {
		logx.Errorf("FindByMobile mobile: %s error: %v", in.Mobile, err)
		return nil, err
	}
	if u == nil {
		return &user.FindByMobileResponse{}, nil
	}
	return &user.FindByMobileResponse{
		UserId:   int64(u.Id),
		Username: u.Username,
		Avatar:   u.Avatar,
	}, nil
}

package logic

import (
	"context"

	"zhihu/app/user_rpc/internal/svc"
	"zhihu/app/user_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByIdLogic {
	return &FindByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindByIdLogic) FindById(in *user.FindByIdRequest) (*user.FindByIdResponse, error) {
	u, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		logx.Errorf("FindById userId: %d error: %v", in.UserId, err)
		return nil, err
	}

	return &user.FindByIdResponse{
		UserId:   int64(u.Id),
		Username: u.Username,
		Mobile:   u.Mobile,
		Avatar:   u.Avatar,
	}, nil
}

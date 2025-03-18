package logic

import (
	"context"
	"time"

	"zhihu/app/user_rpc/internal/code"
	"zhihu/app/user_rpc/internal/model"
	"zhihu/app/user_rpc/internal/svc"
	"zhihu/app/user_rpc/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	if len(in.Username) == 0 {
		return nil, code.RegisterNameEmpty
	}

	res, err := l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Username: in.Username,
		Mobile:   in.Mobile,
		Avatar:   in.Avatar,
		Ctime:    time.Now(),
		Mtime:    time.Now(),
	})
	if err != nil {
		logx.Errorf("register req: %v error: %v", in, err)
		return nil, err
	}
	userId, err := res.LastInsertId()
	if err != nil {
		logx.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	return &user.RegisterResponse{
		UserId: userId,
		}, nil
}

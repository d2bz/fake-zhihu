package logic

import (
	"context"

	"zhihu/app/like/rpc/internal/svc"
	"zhihu/app/like/rpc/like_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IsThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIsThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IsThumbupLogic {
	return &IsThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IsThumbupLogic) IsThumbup(in *like_pb.IsThumbupRequest) (*like_pb.IsThumbupResponse, error) {
	// todo: add your logic here and delete this line

	return &like_pb.IsThumbupResponse{}, nil
}

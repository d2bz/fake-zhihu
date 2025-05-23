package logic

import (
	"context"

	"zhihu/app/like/rpc/internal/svc"
	"zhihu/app/like/rpc/like_pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ThumbupLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewThumbupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ThumbupLogic {
	return &ThumbupLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ThumbupLogic) Thumbup(in *like_pb.ThumbupRequest) (*like_pb.ThumbupResponse, error) {
	// todo: add your logic here and delete this line

	return &like_pb.ThumbupResponse{}, nil
}

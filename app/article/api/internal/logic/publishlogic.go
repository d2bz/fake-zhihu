package logic

import (
	"context"
	"encoding/json"

	"zhihu/app/article/api/internal/code"
	"zhihu/app/article/api/internal/svc"
	"zhihu/app/article/api/internal/types"
	"zhihu/app/article/rpc/article"
	"zhihu/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

const minContentLen = 10

type PublishLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishLogic) Publish(req *types.PublishRequest) (resp *types.PublishResponse, err error) {
	if len(req.Title) == 0 {
		return nil, code.ArticleTitleEmpty
	}
	if len(req.Content) < minContentLen {
		return nil, code.ArticleContentTooFewWords
	}
	userId, err := l.ctx.Value("userId").(json.Number).Int64()
	if err != nil {
		logx.Errorf("get userId from l.ctx error: %v", err)
		return nil, xcode.NoLogin
	}
	pres, err := l.svcCtx.ArticleRPC.Publish(l.ctx, &article.PublishRequest{
		UserId:      userId,
		Title:       req.Title,
		Content:     req.Content,
		Description: req.Description,
		Cover:       req.Cover,
	})
	if err != nil {
		logx.Errorf("l.svcCtx.ArticleRPC.Publish req: %v error: %v", req, err)
		return nil, err
	}
	return &types.PublishResponse{
		ArticleId: pres.ArticleId,
	}, nil
}

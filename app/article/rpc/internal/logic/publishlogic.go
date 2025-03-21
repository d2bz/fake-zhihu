package logic

import (
	"context"

	"zhihu/app/article/rpc/article_pb"
	"zhihu/app/article/rpc/internal/model"
	"zhihu/app/article/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPublishLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishLogic {
	return &PublishLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PublishLogic) Publish(in *article_pb.PublishRequest) (*article_pb.PublishResponse, error) {
	res, err := l.svcCtx.ArticleModel.Insert(l.ctx, &model.Article{
		AuthorId:    uint64(in.UserId),
		Title:       in.Title,
		Content:     in.Content,
		Description: in.Description,
		Cover:       in.Cover,
	})
	if err != nil {
		l.Logger.Errorf("Publish Insert req: %v error: %v", in, err)
		return nil, err
	}
	articleId, err := res.LastInsertId()
	if err != nil {
		l.Logger.Errorf("LastInsertId error: %v", err)
		return nil, err
	}
	return &article_pb.PublishResponse{
		ArticleId: articleId,
	}, nil
}

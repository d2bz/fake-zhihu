// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: article.proto

package server

import (
	"context"

	"zhihu/app/article/rpc/article_pb"
	"zhihu/app/article/rpc/internal/logic"
	"zhihu/app/article/rpc/internal/svc"
)

type ArticleServer struct {
	svcCtx *svc.ServiceContext
	article_pb.UnimplementedArticleServer
}

func NewArticleServer(svcCtx *svc.ServiceContext) *ArticleServer {
	return &ArticleServer{
		svcCtx: svcCtx,
	}
}

func (s *ArticleServer) Publish(ctx context.Context, in *article_pb.PublishRequest) (*article_pb.PublishResponse, error) {
	l := logic.NewPublishLogic(ctx, s.svcCtx)
	return l.Publish(in)
}

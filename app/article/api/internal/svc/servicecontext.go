package svc

import (
	"zhihu/app/article/api/internal/config"
	"zhihu/app/article/rpc/article"
	"zhihu/pkg/interceptors"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	ArticleRPC article.Article
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	articleRPC := zrpc.MustNewClient(c.ArticleRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	return &ServiceContext{
		Config:     c,
		ArticleRPC: article.NewArticle(articleRPC),
	}
}

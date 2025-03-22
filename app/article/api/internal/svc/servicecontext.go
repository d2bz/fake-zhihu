package svc

import (
	"zhihu/app/article/api/internal/config"
	"zhihu/app/article/rpc/article"
	"zhihu/pkg/interceptors"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/zeromicro/go-zero/zrpc"
)

const (
	defaultOssConnectTimeout   = 1
	defaultOssReadWriteTimeout = 3
)

type ServiceContext struct {
	Config     config.Config
	ArticleRPC article.Article
	OssClient  *oss.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 自定义拦截器
	articleRPC := zrpc.MustNewClient(c.ArticleRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))

	if c.Oss.ConnectTimeout == 0 {
		c.Oss.ConnectTimeout = defaultOssConnectTimeout
	}
	if c.Oss.ReadWriteTimeout == 0 {
		c.Oss.ReadWriteTimeout = defaultOssReadWriteTimeout
	}
	oc, err := oss.New(c.Oss.Endpoint, c.Oss.AccessKeyId, c.Oss.AccessKeySecret, oss.Timeout(c.Oss.ConnectTimeout, c.Oss.ReadWriteTimeout))
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:     c,
		ArticleRPC: article.NewArticle(articleRPC),
		OssClient:  oc,
	}
}

package svc

import (
	"zhihu/app/article/rpc/internal/config"
	"zhihu/app/article/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	ArticleModel model.ArticleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSourse)

	return &ServiceContext{
		Config: c,
		ArticleModel: model.NewArticleModel(conn),
	}
}

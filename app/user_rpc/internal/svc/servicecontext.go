package svc

import (
	"zhihu/app/user_rpc/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config   config.Config
	MysqlCon sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		MysqlCon: sqlx.NewMysql(c.Mysql.DataSource),
	}
}

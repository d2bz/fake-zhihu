package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Mysql mysql
}

type mysql struct {
	DataSourse string
}
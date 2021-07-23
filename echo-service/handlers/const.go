package handlers

import (
	"core/echo-service/svc"
	"github.com/go-kit/kit/sd/etcdv3"
)

var (
	// Register 注册中心
	Register *etcdv3.Registrar
	// 配置文件
	globalConfig svc.Config
)

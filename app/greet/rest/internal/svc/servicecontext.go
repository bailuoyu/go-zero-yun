package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-yun/app/greet/rest/internal/config"
	"go-zero-yun/app/greet/rest/internal/middleware"
)

type ServiceContext struct {
	Config         config.Config
	AuthMiddleware rest.Middleware
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		AuthMiddleware: middleware.NewAuthMiddleware().Handle,
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}

package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-yun/app/greet/rest/internal/handler/base"
	"go-zero-yun/app/greet/rest/internal/svc"
	"net/http"
)

func OldRegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/old/ping",
				Handler: base.PingHandler(serverCtx),
			},
		},
	)
}

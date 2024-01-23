package base

import (
	"go-zero-yun/app/greet/rest/internal/logic/base"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-yun/public/handler"
)

func PingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := utils.NewElapsedTimer()

		var req types.Empty
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		r = handler.Request(r, req)

		l := base.NewPingLogic(r.Context(), svcCtx)
		resp, err := l.Ping(&req)
		handler.Response(w, r, resp, err, timer)
	}
}

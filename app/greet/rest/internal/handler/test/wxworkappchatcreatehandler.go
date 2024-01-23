package test

import (
	"net/http"

	"github.com/zeromicro/go-zero/core/utils"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-yun/app/greet/rest/internal/logic/test"
	"go-zero-yun/app/greet/rest/internal/svc"
	"go-zero-yun/app/greet/rest/internal/types"
	"go-zero-yun/public/handler"
)

func WxWorkAppChatCreateHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		timer := utils.NewElapsedTimer()

		var req types.TestWxWorkAppChatCreateReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		r = handler.Request(r, req)

		l := test.NewWxWorkAppChatCreateLogic(r.Context(), svcCtx)
		resp, err := l.WxWorkAppChatCreate(&req)
		handler.Response(w, r, resp, err, timer)
	}
}

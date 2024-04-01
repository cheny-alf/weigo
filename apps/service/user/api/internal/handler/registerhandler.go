package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
	"wego/apps/service/user/api/internal/logic"
	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register(r.Context(), &req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

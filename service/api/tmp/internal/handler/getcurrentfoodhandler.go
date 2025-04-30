package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"menul-service/service/api/tmp/internal/logic"
	"menul-service/service/api/tmp/internal/svc"
	"menul-service/service/api/tmp/internal/types"
)

func GetCurrentFoodHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetCurrentFoodReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewGetCurrentFoodLogic(r.Context(), svcCtx)
		resp, err := l.GetCurrentFood(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

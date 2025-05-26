package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"menul-service/service/api/internal/logic"
	"menul-service/service/api/internal/svc"
)

func GetOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetOrderLogic(r.Context(), svcCtx)
		resp, err := l.GetOrder()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"menul-service/service/api/internal/logic"
	"menul-service/service/api/internal/svc"
)

func GetFoodCategoryListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetFoodCategoryListLogic(r.Context(), svcCtx)
		resp, err := l.GetFoodCategoryList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}

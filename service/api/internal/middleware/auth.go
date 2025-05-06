package middleware

import (
	"errors"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"menul-service/service/cache"
	"net/http"
	"strings"
)

func Auth() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {

		return func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			token = strings.TrimPrefix(token, "Bearer ")
			isPass, err := ValidateToken(token)
			if err != nil {
				httpx.Error(w, err)
			}
			if token == "" || !isPass {
				httpx.Error(w, errors.New("unauthorized"))
				return
			}

			next(w, r)
		}
	}
}

func ValidateToken(token string) (bool, error) {
	getUserId, err := cache.GetToken(token)
	if err != nil {
		return false, err
	}
	if getUserId == "" {
		return false, nil
	}
	return true, nil

}

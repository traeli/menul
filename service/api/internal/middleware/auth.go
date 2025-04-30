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
			userID := r.Header.Get("user_id")

			token = strings.TrimPrefix(token, "Bearer ")
			isPass, err := ValidateToken(token, userID)
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

func ValidateToken(token string, userID string) (bool, error) {
	getToken, err := cache.GetToken(userID)
	if err != nil {
		return false, err
	}
	if getToken == token {
		return true, nil
	}
	return false, nil

}

package middleware

import (
	"errors"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"menul-service/service/cache"
	"net/http"
	"strings"
	"time"
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

func GetTimePeriod(t time.Time) string {
	hour := t.Hour()

	switch {
	case hour >= 5 && hour < 12:
		return "breakfast"
	case hour >= 12 && hour < 15:
		return "lunch"
	case hour >= 15 && hour < 18:
		return "dinner"
	default:
		return "all"
	}

}

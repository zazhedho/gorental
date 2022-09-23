package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/zazhedho/gorental/src/helpers"
)

func CheckAuth(role string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")

		if !strings.Contains(headerToken, "Bearer") {
			helpers.New("invalid header type", 401, true).Send(w)
			return
		}
		token := strings.Replace(headerToken, "Bearer ", "", -1)

		checkTokens, err := helpers.CheckToken(token, role)
		if err != nil {
			helpers.New(err.Error(), 401, true).Send(w)
			return
		}

		if !checkTokens {
			helpers.New("anda bukan admin", 403, true).Send(w)
			return
		}

		eksportToken, err := helpers.EksportToken(token)
		if err != nil {
			helpers.New(err.Error(), 401, true).Send(w)
			return
		}

		ctx := context.WithValue(r.Context(), "username", eksportToken.Username)

		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

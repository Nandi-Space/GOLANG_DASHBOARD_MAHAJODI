package api

import (
	"Mahajodi_GOLANG_Dashboard/store"
	"Mahajodi_GOLANG_Dashboard/utils"
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/sirupsen/logrus"
)

func Authentication(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if len(token) > 7 {
			token = token[7:]
		}

		parsedToken, err := utils.ParseJWTToken(token)
		if err != nil {
			logrus.Error(err)
			utils.ErrorResponse(w, "failed to verify token", 401)
			return
		}
		admin, err := store.DBState.GetAdmin2(parsedToken.UserID)
		if err != nil {
			logrus.Error("invalid admin id in jwt token, ", err)
			utils.ErrorResponse(w, "invalid admin_id", 500)
			return
		}
		ctx := context.WithValue(r.Context(), "admin", admin)
		next(w, r.WithContext(ctx))
	})

}

func DumpRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dump, err := httputil.DumpRequest(r, true)
		if err != nil {
			logrus.Error(err)
		}
		fmt.Println(string(dump))
		next.ServeHTTP(w, r)
	})
}

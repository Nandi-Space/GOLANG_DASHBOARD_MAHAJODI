package api

import (
	"log"
	"mahajodi/dashboard/handlers/auth"
	"mahajodi/dashboard/handlers/count"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func routes() *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1/dashboard", func(r chi.Router) {
		r.Get("/totalmembers", count.GetTotalMembers)
		r.Get("/totalmale", count.GetTotalMales)
		r.Get("/totalfemale", count.GetTotalFemales)
		r.Get("/data-male", count.GetMales)
		r.Get("/data-female", count.GetFemales)
		r.Post("/login-with-mobile", auth.VerifyEmail)
		r.Post("/login-with-mobile/verify-otp", auth.VerifyOtp)
	})
	return r
}

func StartServer(listenAddr string) {
	r := routes()
	log.Println("server started at:", listenAddr)
	logrus.Error(http.ListenAndServe(listenAddr, r))
}

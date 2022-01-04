package api

import (
	"Mahajodi_GOLANG_Dashboard/handlers/auth"
	"Mahajodi_GOLANG_Dashboard/handlers/count"
	"Mahajodi_GOLANG_Dashboard/handlers/payment"
	"Mahajodi_GOLANG_Dashboard/handlers/users"
	"log"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func routes() *chi.Mux {
	r := chi.NewRouter()

	//Uncomment this to dump all API requests
	//r.Use(DumpRequest)

	r.Route("/api/v1/dashboard", func(r chi.Router) {

		r.Get("/totalmembers", Authentication(count.GetTotalMembers))
		r.Get("/totalmale", Authentication(count.GetTotalMales))
		r.Get("/totalfemale", Authentication(count.GetTotalFemales))

		r.Get("/payment-details", Authentication(payment.GetPayment))

		r.Get("/data-male", Authentication(users.GetMales))
		r.Get("/data-female", Authentication(users.GetFemales))
		r.Get("/new-users", Authentication(users.GetNewUsers))

		r.Post("/login-with-email", auth.VerifyEmail)
		r.Post("/login-with-email/verify-otp", auth.VerifyOtp)
	})

	//Deprecated old APIs
	r.Route("/api/v0/", func(r chi.Router) {

	})

	return r
}

func StartServer(listenAddr string) {
	r := routes()

	log.Println("server started at:", listenAddr)

	logrus.Error(http.ListenAndServe(listenAddr, r))
}

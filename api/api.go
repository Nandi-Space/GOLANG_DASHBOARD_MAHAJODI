package api

import (
	"fmt"

	"net/http"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
)

func routes() *chi.Mux {
	r := chi.NewRouter()

	//Uncomment this to dump all API requests
	//r.Use(DumpRequest)

	r.Route("/api/v1/", func(r chi.Router) {

	})

	//Deprecated old APIs
	r.Route("/api/v0/", func(r chi.Router) {

	})

	return r
}

func StartServer(listenAddr string) {
	r := routes()
	fmt.Println("started server at:", listenAddr)
	logrus.Error(http.ListenAndServe(listenAddr, r))
}

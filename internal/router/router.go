package router

import (
	"log"
	"net/http"

	"github.com/alcbotta/go-prometheus-grafana/tree/develop/internal/handler/option"
	"github.com/alcbotta/go-prometheus-grafana/tree/develop/internal/handler/user"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer() {
	r := mux.NewRouter()

	r.Handle("/metrics", promhttp.Handler()).Methods("GET")

	r.HandleFunc("/user/{id}", user.GetUserHandler).Methods("GET")
	r.HandleFunc("/user", user.PostUserHandler).Methods("POST")

	// Rota OPTIONS para responder a requisições preflight
	r.HandleFunc("/user/{id}", option.OptionsHandler).Methods("OPTIONS")
	r.HandleFunc("/user", option.OptionsHandler).Methods("OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", r))
}

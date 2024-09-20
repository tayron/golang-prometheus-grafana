package user

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/alcbotta/go-prometheus-grafana/tree/develop/internal/metric"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var metricUserGet *prometheus.CounterVec

func init() {
	metricUserGet = metric.NewUserGetStatusMetric()
	prometheus.MustRegister(metricUserGet)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	var status string

	var requestData struct {
		User string
	}

	json.NewDecoder(r.Body).Decode(&requestData)

	if rand.Float32() > 0.8 {
		status = "4xx"
		log.Printf("user #%s userID not found ", userID)
	} else {
		status = "2xx"
		log.Printf("user #%s userID found ", userID)
	}

	metricUserGet.WithLabelValues(requestData.User, status).Inc()

	w.Write([]byte(status))
}

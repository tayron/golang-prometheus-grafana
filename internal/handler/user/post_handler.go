package user

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/alcbotta/go-prometheus-grafana/tree/develop/internal/metric"
	"github.com/prometheus/client_golang/prometheus"
)

var metricUserPost *prometheus.CounterVec

func init() {
	metricUserPost = metric.NewUserPostStatusMetric()
	prometheus.MustRegister(metricUserPost)
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var status string

	var requestData struct {
		User string
	}
	json.NewDecoder(r.Body).Decode(&requestData)

	if rand.Float32() > 0.8 {
		status = "user_not_created"
		log.Println("user not created")
	} else {
		status = "user_created"
		log.Println("user created")
	}

	metricUserPost.WithLabelValues(status).Inc()

	w.Write([]byte(status))
}

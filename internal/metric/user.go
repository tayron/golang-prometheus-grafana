package metric

import "github.com/prometheus/client_golang/prometheus"

func NewUserGetStatusMetric() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_get_user_status_count", // metric name
			Help: "Conta o número de status http retornado ao realizar o get em usuários.",
		},
		[]string{"user", "status"}, // labels
	)
}

func NewUserPostStatusMetric() *prometheus.CounterVec {
	return prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_request_post_user_status_count", // metric name
			Help: "Conta o númemro de status http retornado na hora de criar um usuário.",
		},
		[]string{"status"}, // labels
	)
}

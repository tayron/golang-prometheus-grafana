package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/alcbotta/go-prometheus-grafana/tree/develop/internal/router"
)

type MyRequest struct {
	User string
}

func producerUserRequest() {
	userPool := []string{"bob", "alice", "jack"}
	for {
		postBody, _ := json.Marshal(MyRequest{
			User: userPool[rand.Intn(len(userPool))],
		})
		requestBody := bytes.NewBuffer(postBody)

		// Get user with ID 1
		http.Get("http://localhost:8080/user/1")
		time.Sleep(time.Second * 2)

		// Create new user
		http.Post("http://localhost:8080/user", "application/json", requestBody)
		time.Sleep(time.Second * 2)
	}
}

func main() {
	go producerUserRequest()
	router.StartServer()
}

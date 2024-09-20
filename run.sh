
#!/bin/bash

go build -o web_application -ldflags "-s -w" cmd/main.go
docker-compose down
docker-compose up --build -d
docker logs server -f


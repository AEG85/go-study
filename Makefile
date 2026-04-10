include .env
export

start-service:
	@export NEW_USER=YES && \
	go run main.go
include .env
export

run-http-app:
	@go run main.go

postgres-up:
	@docker-compose up \
		-d \
		postgres
http-service-up:
	@docker-compose up \
		-d \
		application
run-docker-http-app:
	@docker-compose up -d
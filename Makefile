include .env
export

run-docker:
	@docker run \
	-d \
	--rm \
	--name test-docker \
	-p 9091:9091 \
	hw-docker-3
run-http-app:
	@go run main.go
postgres-up:
	@docker run \
		-d \
		--rm \
		-p 5432:5432 \
		-e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
		-v .out/pgdata:/var/lib/postgresql \
		postgres:18.3-bookworm
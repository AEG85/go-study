run-docker:
	@docker run \
	-d \
	--rm \
	--name test-docker \
	-p 9091:9091 \
	hw-docker-3
run-http-app:
	@go run main.go
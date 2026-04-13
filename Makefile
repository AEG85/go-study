run-http-app:
	@docker run \
	-d \
	--rm \
	--name test-docker \
	-p 9091:9091 \
	hw-docker-3
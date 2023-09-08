.PHONY: docker
docker:
	@rm home3 || true
	@GOOS=linux GOARCH=arm go build -o home3 .
	@docker rmi -f lee/home3:v0.0.1
	@docker build -t lee/home3-live:v0.0.1 .
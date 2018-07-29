build:
	GOOS=linux GOARCH=amd64 go build
	#docker rmi service
	docker build -t service:v1 .
run:
	docker run --name="test_service" -e REGISTER_ADDR="192.168.3.28:50001" -p 50001:50001 service
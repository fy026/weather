build:
	GOOS=linux GOARCH=amd64 go build
	docker rmi gateway
	docker build -t gateway:latest .
run:
	docker run --name="gateway" -p 8001:8001 gateway 

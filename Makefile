proto:
	protoc pkg/pb/*.proto --go_out=plugins=grpc:.

chat-service:
	go run cmd/main.go

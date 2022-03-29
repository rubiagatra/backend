dev-api:
	go run service/api/main.go serve

migrate-api:
	go run service/api/main.go migrate

proto:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    service/api/pb/helloworld/helloworld.proto

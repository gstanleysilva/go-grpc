protoc:
	@protoc --go_out=. --go-grpc_out=. proto/course_category.proto

run_grpc:
	@go run cmd/grpcServer/main.go

evans:
	@evans -r repl
.PHONY: all generate tidy build

generate_api:
	protoc --go_out=api/proto \
	--go_opt=module=github.com/fmuyassarov/grpc-go  \
	--go-grpc_out=api/proto \
	--go-grpc_opt=module=github.com/fmuyassarov/grpc-go \
	api/proto/api.proto

generate_cal:
	protoc --go_out=calculator/proto \
	--go_opt=module=github.com/fmuyassarov/grpc-go  \
	--go-grpc_out=calculator/proto \
	--go-grpc_opt=module=github.com/fmuyassarov/grpc-go \
	calculator/proto/calculator.proto

tidy:
	go mod tidy

build_api: 
	go build -o cmd/api/client ./api/client
	go build -o cmd/api/server ./api/server

build_cal: 
	go build -o cmd/cal/client ./calculator/client
	go build -o cmd/cal/server ./calculator/server

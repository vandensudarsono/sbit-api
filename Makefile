BIN_NAME=sbit

dep:
	@echo "> fetching dependencies..."
	@go mod tidy

install: dep
	@echo "> building binary..."
	@CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' -o ./${BIN_NAME} main.go


run: dep install
	@echo "> running application..."
	@./${BIN_NAME} serve

protoc-gen:
	@echo "> generating protobuf files..."
	@cd ./shared/proto/wallet && protoc *.proto -I. --go_out=../../../infrastructure/grpc/proto/wallet  --go-grpc_out=../../../infrastructure/grpc/proto/wallet

protoc-gen-clean:
	@echo "> cleaning protobuf files..."
	@cd ./infrastructure/grpc/proto/validator && rm -rf *.pb.go
	@cd ./infrastructure/grpc/proto/wallet && rm -rf *.pb.go

eal.PHONY: dep install run protoc-gen protoc-gen-clean setup
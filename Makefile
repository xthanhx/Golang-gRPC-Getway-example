TARGET=helloworld



all: build-proto



clean:

	rm -rf $(TARGET)


build:

	go build -o $(TARGET) main.go


build-proto:

	protoc -I ./proto \
   --go_out ./proto --go_opt paths=source_relative \
   --go-grpc_out ./proto --go-grpc_opt paths=source_relative \
   --grpc-gateway_out ./proto --grpc-gateway_opt paths=source_relative \
   ./proto/helloworld/helloworld.proto

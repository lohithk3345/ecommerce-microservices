all: protoc-user protoc-product build

build: userService productService orderService cartService
	# @go build -o bin/ecommerce use_main.go

clean:
	@rm -rf bin/

build-run: build
	@./bin/userService

runner:
	@go run ./cmd/userService

run:

run-userService:
	@./bin/userService

run-productService:
	@./bin/productService

run-orderService:
	@./bin/orderService

run-cartService:
	@./bin/cartService

userService:
	@go build -o bin/userService ./cmd/userService

productService:
	@go build -o bin/productService ./cmd/productService

orderService:
	@go build -o bin/orderService ./cmd/orderService

cartService:
	@go build -o bin/cartService ./cmd/cartService

protoc-user:
	@protoc --go_out=buffers/userpb/ --go_opt=paths=source_relative \
    --go-grpc_out=buffers/userpb/ --go-grpc_opt=paths=source_relative \
    protobuffs/user.proto

protoc-product:
	@protoc --go_out=buffers/productpb --go_opt=paths=source_relative \
    --go-grpc_out=buffers/productpb/ --go-grpc_opt=paths=source_relative \
    protobuffs/product.proto

protoc-clean:
	@rm -rf buffers/usepb/*
	@rm -rf buffers/productpb/*

protoc-all: protoc-user protoc-product

all: build

build:
	@go build -o bin/ecommerce ./cmd/userService

clean:
	@rm -rf bin/

build-run: build
	@./bin/ecommerce

runner:
	@go run ./cmd/userService

run:
	@./bin/ecommerce

userService:
	@go build -o bin/userService ./cmd/userService

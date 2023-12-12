all: build

build:
	@go build -o bin/ecommerce main.go

clean:
	@rm -rf bin/

build-run: build
	@./bin/ecommerce

runner:
	@go run main.go

run:
	@./bin/ecommerce

userService:
	@go build -o bin/userService ./cmd/userService

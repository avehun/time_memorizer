all: clean build

build: 
	go build -o cmd/time_memorizer/time_memorizer -v ./cmd/time_memorizer

run:
	go build -o cmd/time_memorizer/time_memorizer -v ./cmd/time_memorizer
	./cmd/time_memorizer/time_memorizer

clean: 
	go clean
	rm cmd/time_memorizer/time_memorizer

lint:
	golangci-lint run ./...

protos:
	protoc -I api/protos/ --go_out=pkg/api/ --go-grpc_out=pkg/api api/protos/timeMemorizer.proto

local_up:
	docker-compose up --build

ci: build lint

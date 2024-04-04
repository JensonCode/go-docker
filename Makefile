build:
	@go build -o bin/go-docker cmd/go-docker/main.go

run: build
	@./bin/go-docker

test:
	@go test -v ./...
	
docker-run-local:
	@docker-compose up --build
build:
	@go build -o bin/go-docker main.go

run: build
	@./bin/go-docker

test:
	@go test -v ./...
	
docker-run-local:
	@docker-compose up --build
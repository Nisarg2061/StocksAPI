build:
	@go build -o bin/fs 

run: 
	@go run ./cmd/main.go

test:
	@go test ./... -v

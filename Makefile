default:
	@echo "have some Makefile"

run-test:
	go run -race cmd/test.go

run-api:
	go run -race cmd/api.go

test:
	go test -race -v ./...

build-test:
	go build -race cmd/test.go
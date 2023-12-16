run:
	go run main.go

build:
	go build -o tfdocs

mod-vendor:
	go mod vendor

linter:
	@golangci-lint run

gosec:
	@gosec -quiet ./...

validate: linter gosec
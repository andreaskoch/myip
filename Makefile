build:
	go build
	go test -cover

test:
	go test

coverage:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

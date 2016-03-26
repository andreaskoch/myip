build:
	go test -tags='!integration' -cover
	go build

test:
	go test -tags='!integration'

integrationtest:
	go test -tags='integration'

coverage:
	go test -tags='!integration' -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

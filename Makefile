.PHONY: test
test: 
	go test -v -cover ./...

.PHONY: cover
cover:
	go test -coverprofile=coverage.out ./...
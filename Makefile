.PHONY: test
test:
	@echo "running tests"
	go test -v ./...

.PHONY: clean
clean:
	@echo "cleaning releases"
	@GOOS=linux go clean -i -x ./...
	@rm -rf build/

.PHONY: generate
generate:
	go install github.com/swaggo/swag/cmd/swag@latest
	go generate ./...
	go mod tidy

.PHONY: swagger
swagger:
	swag init -g ./cmd/server/main.go -o /docs/swagger --parseDependency

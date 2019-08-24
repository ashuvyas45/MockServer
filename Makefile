APP_EXECUTABLE="out/mock-server"
UNIT_TEST_PACKAGES=$(shell go list ./... | grep -v "vendor")
setup:
    go get golang.org/x/lint/golint

compile:
    mkdir -p out/
    go build -o $(APP_EXECUTABLE)

build: compile fmt vet lint

fmt:
	go fmt ./...

vet:
	go vet ./...

lint:
	@for p in $(UNIT_TEST_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; } \
	done

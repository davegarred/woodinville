GO_GET := go get -u
GO_CLEAN := go clean
GO_TEST := go test
GO_BUILD := go build
TARGET := wine

.PHONY: deps
deps:
	$(GO_GET) github.com/stretchr/testify/assert
	$(GO_GET) github.com/looplab/eventhorizon

.PHONY: clean
clean:
	$(GO_CLEAN)
	rm -f $(TARGET)

.PHONY: test
test:
		$(GO_TEST) -v -short -cover github.com/davegarred/woodinville...


.PHONY: build
build: test
		$(GO_BUILD) -v -gcflags "-N -l" -o $(TARGET)

.PHONY: docker
docker:
		docker build -t wine .
		docker run --rm -p 8000:8000 wine
GO_TEST := go test
GO_BUILD := go build
TARGET := wine

.PHONY: deps
deps:
	go get -u github.com/stretchr/testify/assert

.PHONY: clean
clean:
	go clean
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
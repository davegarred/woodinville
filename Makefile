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

.PHONY: cover
cover:
		go test -short -coverprofile=.coverage github.com/davegarred/woodinville...
		go tool cover -html=.coverage

.PHONY: build
build: test
		$(GO_BUILD) -v -gcflags "-N -l" -o $(TARGET)

.PHONY: package
package: build
		docker build -t wine .
		docker save -o wine_image.tar wine

.PHONY: docker
docker: package
		docker run \
			-d \
			--name wine \
			-p "8500:8500" \
			--rm \
			wine
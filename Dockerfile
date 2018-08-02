FROM golang:1.8

WORKDIR /go/src/app
COPY . .

RUN go get -d -v github.com/stretchr/testify/assert
RUN go install -v ./...

CMD ["app"]
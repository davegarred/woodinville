FROM golang:1.8

WORKDIR /app

# Set an env var that matches your github repo name, replace treeder/dockergo here with your repo name
ENV SRC_DIR=/go/src/github.com/davegarred/woodinville/
# Add the source code:
ADD . $SRC_DIR

RUN cd $SRC_DIR; go build -o woodinville; cp woodinville /app/
ENTRYPOINT ["./woodinville"]
#CMD ["./woodinville"]
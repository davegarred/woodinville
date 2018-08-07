package main

import (
	"os"
	"os/signal"
	"github.com/davegarred/woodinville/web"
)

func main() {
	s := web.Serve()
	waitForSigint()
	s.Shutdown()
}

func waitForSigint() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}


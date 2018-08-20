package main

import (
	"os"
	"os/signal"
	"github.com/davegarred/woodinville/web"
)

var key = "AIzaSyC4R6AN7SmujjPUIGKdyao2Kqitzr1kiRg"

func main() {
	s := web.Serve(key)
	waitForSigint()
	s.Shutdown()
}

func waitForSigint() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}


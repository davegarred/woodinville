package main

import (
	"os"
	"os/signal"
)

func main() {
	s := Serve()
	waitForSigint()
	s.Shutdown()
}

func waitForSigint() {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)
	<-sigint
}


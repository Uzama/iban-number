package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Uzama/iban-number/http"
)

func main() {

	// init routers
	r := http.InitRouter()

	// start the server
	http.StartServer(r)

	channel := make(chan os.Signal, 1)

	signal.Notify(channel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	// wait till to recive a signal to stop the server
	<-channel

	os.Exit(0)
}

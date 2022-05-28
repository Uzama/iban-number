package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Uzama/iban-number/http"
)

func main() {

	r := http.InitRouter()

	http.StartServer(r)

	channel := make(chan os.Signal, 1)

	signal.Notify(channel, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)

	<-channel

	os.Exit(0)
}

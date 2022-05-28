package http

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// start a http server at 0.0.0.0:8080
func StartServer(router *mux.Router) {

	address := "0.0.0.0:8080"

	// create a server
	server := &http.Server{
		Addr:         address,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 10,
		Handler:      router,
	}

	// run the server without blocking
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Println(err.Error())
		}
	}()

	log.Printf("server started at %s", address)
}

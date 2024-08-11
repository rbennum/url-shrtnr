package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// TODO: init config
	// TODO: establish DB connection

	gin_handler := gin.Default()

	serv := &http.Server{
		Addr: "",
		Handler: gin_handler,
	}

	go func() {
		log.Printf("Listening to %s:%s...", "addr_route", "port")
		if err := serv.ListenAndServe(); 
				err != nil && err != http.ErrServerClosed {
				log.Fatalf("Listen: %s\n", err)
		}
}()
}
package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/config"
)

func main() {
	config.LoadConfig()
	// TODO: establish DB connection

	// gin_handler := gin.Default()

	// serv := createServer("", gin_handler)

}

func createServer(addr string, handler *gin.Engine) *http.Server {
	serv := &http.Server{
		Addr: addr,
		Handler: handler,
	}
	
	go func() {
		log.Printf("Listening to %s:%s...", "addr_route", "port")
		if err := serv.ListenAndServe(); 
				err != nil && err != http.ErrServerClosed {
				log.Fatalf("Listen: %s\n", err)
		}
	}()

	return serv
}
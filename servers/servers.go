package servers

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/db"
)

func CreateServer(addr string, handler *gin.Engine) *http.Server {
	serv := &http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		log.Printf("Listening to %s...", addr)
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Error listening to server: %s\n", err)
		}
	}()
	return serv
}

func InitiateShutdown(serv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Printf("Server %s exiting...", serv.Addr)
	db.Pool_DB.Close()
}

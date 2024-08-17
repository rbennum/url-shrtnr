package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/db"
	"github.com/rbennum/url-shrtnr/repositories"
	"github.com/rbennum/url-shrtnr/routes"
	"github.com/rbennum/url-shrtnr/services"
	"github.com/rbennum/url-shrtnr/utils"
)

func main() {
	utils.Init()

	// initiate logger
	err := utils.InitLogger()
	if err != nil {
		log.Fatal("Unable to init logger:", err)
	}

	// load config from .env file
	config, err := utils.LoadConfig()
	if err != nil {
		log.Fatal("Unable to init config:", err)
	}

	// initiate DB connection
	connectDB(&config)

	// initiate both servers
	main_handler := configureMainHandler()
	main_serv := createServer(
		fmt.Sprintf("%s:%s", config.MainServerAddr, config.MainServerPort),
		main_handler,
	)
	short_handler := configureShortHandler()
	short_serv := createServer(
		fmt.Sprintf("%s:%s", config.ShortServerAddr, config.ShortServerPort),
		short_handler,
	)

	// initiate shutdown if triggered
	initiateShutdown(main_serv)
	initiateShutdown(short_serv)
}

func createServer(addr string, handler *gin.Engine) *http.Server {
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

func connectDB(config *utils.CommonConfig) {
	errdb := db.Init(config)
	if errdb != nil {
		log.Fatal("Unable to init DB:", errdb)
	}
}

func initiateShutdown(serv *http.Server) {
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

func configureMainHandler() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/main/*")
	r.Static("/static", "./views/static")

	// init short feature
	repo := repositories.NewShortRepository(db.Pool_DB)
	service := services.NewShortService(repo)
	mainRouter := routes.NewMainRoute(&service)
	routes.CreateMainRoute(mainRouter, r)

	return r
}

func configureShortHandler() *gin.Engine {
	r := gin.Default()

	repo := repositories.NewShortRepository(db.Pool_DB)
	service := services.NewShortService(repo)
	shortRouter := routes.NewShortRouter(&service)
	routes.CreateShortRoute(shortRouter, r)
	return r
}

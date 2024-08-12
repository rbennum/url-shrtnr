package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rbennum/url-shrtnr/config"
	"github.com/rbennum/url-shrtnr/db"
)

func main() {
	// load config from .env file
	config.LoadConfig()

	// initiate DB connection
	connectDB()

	gin_handler := gin.Default()
	gin_handler.LoadHTMLGlob("views/*")
	gin_handler.GET("/", func(ctx *gin.Context) {
		ctx.HTML(
			http.StatusOK,
			"main.html",
			gin.H {
				"status": "success",
			},
		)
	})

	main_addr := config.GetEnv("ADDR_ROUTE", "localhost")
	main_port := config.GetEnv("PORT", "8080")
	short_addr := config.GetEnv("ADDR_ROUTE_SHORTEN", "localhost")
	short_port := config.GetEnv("PORT_SHORTEN", "8088")
	main_serv := createServer(main_addr + ":" + main_port, gin_handler)
	short_serv := createServer(short_addr + ":" + short_port, gin_handler)

	// initiate shutdown if triggered
	initiateShutdown(main_serv)
	initiateShutdown(short_serv)
}

func createServer(addr string, handler *gin.Engine) *http.Server {
	serv := &http.Server{
		Addr: addr,
		Handler: handler,
	}
	go func() {
		log.Printf("Listening to %s...", addr)
		if err := serv.ListenAndServe(); 
				err != nil && err != http.ErrServerClosed {
				log.Fatalf("Listen: %s\n", err)
		}
	}()
	return serv
}

func connectDB() {
	port, errconv := strconv.Atoi(config.GetEnv("DB_PORT", "")) 
	if errconv != nil {
		panic(errconv)
	}
	opts := db.PoolOptions {
		Host: config.GetEnv("DB_HOST", ""),
		Port: port,
		User: config.GetEnv("DB_USER", ""),
		DBName: config.GetEnv("DB_NAME", ""),
		Pass: config.GetEnv("DB_PASS", ""),
	}
	errdb := db.Init(opts)
	if errdb != nil {
		panic(errdb)
	}
	migrateDB()
}

func initiateShutdown(serv *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<- quit
	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := serv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Printf("Server %s exiting...", serv.Addr)
	db.Pool_DB.Close()
}

func migrateDB() {
	driver, err := postgres.WithInstance(
		db.Pool_DB.GetInstance().DB,
		&postgres.Config{},
	)
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		panic(err)
	}
	m.Up()
}

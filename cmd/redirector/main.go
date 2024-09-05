package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rbennum/url-shrtnr/db"
	"github.com/rbennum/url-shrtnr/repositories"
	"github.com/rbennum/url-shrtnr/routes"
	"github.com/rbennum/url-shrtnr/servers"
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
	config := utils.LoadConfig()

	// initiate DB connection
	err = db.Init(&config)
	if err != nil {
		log.Fatal("Unable to init DB:", err)
	}

	// initiate the redirector server
	redirector_handler := configureRedirectorHandler(config)
	short_serv := servers.CreateServer(
		fmt.Sprintf("%s:%s", config.ShortServerAddr, config.ShortServerPort),
		redirector_handler,
	)

	// initiate shutdown if triggered
	servers.InitiateShutdown(short_serv)
}

func configureRedirectorHandler(config utils.CommonConfig) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/main/*")
	r.Static("/static", "./views/static")

	repo := repositories.NewShortRepository(db.Pool_DB)
	service := services.NewShortService(repo, config)
	shortRouter := routes.NewShortRouter(&service)
	routes.CreateShortRoute(shortRouter, r)
	return r
}

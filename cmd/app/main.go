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

	// initiate both servers
	main_handler := configureMainHandler(config)
	main_serv := servers.CreateServer(
		fmt.Sprintf("%s:%s", config.MainServerAddr, config.MainServerPort),
		main_handler,
	)

	// initiate shutdown if triggered
	servers.InitiateShutdown(main_serv)
}

func configureMainHandler(config utils.CommonConfig) *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/main/*")
	r.Static("/static", "./views/static")

	// init short feature
	repo := repositories.NewShortRepository(db.Pool_DB)
	service := services.NewShortService(repo, config)
	mainRouter := routes.NewMainRoute(&service)
	routes.CreateMainRoute(mainRouter, r)

	return r
}

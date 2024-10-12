package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/rbennum/url-shrtnr/internals/client"
	"github.com/rbennum/url-shrtnr/internals/db"
	"github.com/rbennum/url-shrtnr/internals/server"
	"github.com/rbennum/url-shrtnr/utils"
	log "github.com/rs/zerolog/log"
)

func main() {
	utils.NewRandomStringSeed()

	config := utils.NewConfig()

	logFn := utils.NewLogger(config)
	defer logFn()

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		oscall := <-ch
		log.Fatal().Msgf("system call: %+v", oscall)
		cancel()
	}()

	err := db.Migrate(config.DBMigrationPath, config.DBSourceURL)
	if err != nil {
		log.Fatal().Err(err).Msgf("Unable to migrate DB: %+v", err)
	}

	server := server.NewServer(
		server.ServerOpts{
			Config: config,
			Client: &client.Client{DB: db.NewDatabase(config.DBSourceName)},
		},
	)
	server.Run(ctx)
}

package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rbennum/url-shrtnr/internals/client"
	"github.com/rbennum/url-shrtnr/internals/middleware"
	"github.com/rbennum/url-shrtnr/internals/repository"
	"github.com/rbennum/url-shrtnr/internals/route"
	"github.com/rbennum/url-shrtnr/internals/service"
	"github.com/rbennum/url-shrtnr/utils"
	"github.com/rs/zerolog/log"
)

type ServerOpts struct {
	Config utils.CommonConfig
	Client *client.Client
}

type Server struct {
	opts       ServerOpts
	urlRepo    *repository.UrlRepo
	urlService service.UrlService
}

func NewServer(opts ServerOpts) Server {
	log.Debug().
		Str("postgres", opts.Config.DBSourceName).
		Msg("checking config")
	s := Server{
		opts: opts,
	}
	s.urlRepo = repository.NewUrlRepo(opts.Client.DB)
	s.urlService = service.NewUrlService(s.urlRepo, &opts.Config)
	return s
}

func (s *Server) Run(ctx context.Context) {
	httpServer := s.newHTTPServer(ctx)
	go func() {
		log.Info().Msgf("Starting shortener service on %s", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msgf("Listen:%+s\n", err)
		}
	}()

	<-ctx.Done()

	gracefulShutdownPeriod := 30 * time.Second

	log.Warn().Msg("Shutting down http server")
	shutdownCtx, cancel := context.WithTimeout(context.Background(), gracefulShutdownPeriod)
	defer cancel()
	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Error().Err(err).Msg("Failed to shutdown http server gracefully")
	}
	log.Warn().Msg("HTTP server gracefully stopped")
}

func (s *Server) newHTTPServer(ctx context.Context) *http.Server {
	mux := mux.NewRouter()

	urlHandler := route.NewUrlHandler(ctx, s.urlService)
	api := mux.PathPrefix("/api/v1").Subrouter()
	api.Path("/url").Handler(urlHandler)
	api.Use(middleware.LoggingMiddleware)
	return &http.Server{
		Addr:    fmt.Sprintf("%s:%s", s.opts.Config.ServerAddr, s.opts.Config.ServerPort),
		Handler: mux,
	}
}

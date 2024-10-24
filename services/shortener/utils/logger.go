package utils

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	log "github.com/rs/zerolog/log"
)

func NewLogger(config CommonConfig) func() {
	level, err := zerolog.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to parse log level")
	}
	zerolog.SetGlobalLevel(level)

	writers := []io.Writer{zerolog.ConsoleWriter{Out: os.Stdout}}

	if err := os.MkdirAll("logs", os.ModePerm); err != nil {
		log.Fatal().Err(err).Msg("Unable to create log directory")
	}

	var runLogFile *os.File
	runLogFile, err = os.OpenFile(
		"logs/log.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to open log file")
	}
	writers = append(writers, runLogFile)

	zerolog.TimeFieldFormat = time.RFC3339Nano

	multi := zerolog.MultiLevelWriter(writers...)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	return func() {
		if runLogFile != nil {
			log.Info().Msg("Closing log file to avoid leak")
			runLogFile.Close()
		}
	}
}

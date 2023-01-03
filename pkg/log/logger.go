package log

import (
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	log     zerolog.Logger
}

func (logger *Logger) Init() {
	logger.log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr,
		TimeFormat: "2006-01-02 15:04:05 MST"}).With().
		Timestamp().
		Logger()
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.log.Error().Msgf(template, args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.log.Info().Msgf(template, args...)
}
func (logger *Logger) Info(msg string) {
	logger.log.Info().Msg(msg)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.log.Fatal().Msgf(template, args...)
}

func (logger *Logger) Fatal(msg string) {
	logger.log.Fatal().Msg(msg)
}

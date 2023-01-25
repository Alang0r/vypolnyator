package log

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

type Logger struct {
	Name string
	log  zerolog.Logger
}

func NewLogger() Logger {
	l := Logger{}
	l.Init()
	return l
}

func (logger *Logger) Init(args...string) {
	if len(args) != 0 {
		logger.Name = args[0]
	}
	logger.log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr,
		TimeFormat: "2006-01-02 15:04:05"}).With().
		Timestamp().
		Logger()
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.log.Error().Msgf(template, args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	tmp := fmt.Sprintf("%s: %s", logger.Name, template)
	logger.log.Info().Msgf(tmp, args...)
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

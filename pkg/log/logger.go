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
	tmp := fmt.Sprintf("%s: %s", logger.Name, template)
	logger.log.Error().Msgf(tmp, args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	tmp := fmt.Sprintf("%s: %s", logger.Name, template)
	logger.log.Info().Msgf(tmp, args...)
}
func (logger *Logger) Info(msg string) {
	tmp := fmt.Sprintf("%s: %s", logger.Name, msg)
	logger.log.Info().Msg(tmp)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	tmp := fmt.Sprintf("%s: %s", logger.Name, template)
	logger.log.Fatal().Msgf(tmp, args...)
}

func (logger *Logger) Fatal(msg string) {
	tmp := fmt.Sprintf("%s: %s", logger.Name, msg)
	logger.log.Fatal().Msg(tmp)
}

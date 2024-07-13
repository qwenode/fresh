package runner

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

type logFunc func(string, ...interface{})

func newLogFunc(prefix string) func(string, ...interface{}) {
	return func(format string, v ...interface{}) {
		log.Info().Str("Type", prefix).Msgf(format, v...)
	}
}

func fatal(err error) {
	log.Fatal().Err(err).Send()
}

type appLogWriter struct{}

func (a appLogWriter) Write(p []byte) (n int, err error) {
	fmt.Print(string(p))
	return len(p), nil
}

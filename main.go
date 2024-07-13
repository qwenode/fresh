package main

import (
	"fresh/runner"
	"github.com/joho/godotenv"
	"github.com/qwenode/gogo/ff"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02T15:04:05"})
	args := os.Args
	if len(args) < 2 || args[1] == "" {
		log.Fatal().Msg("config file is required")
	}
	cfgFile := os.Args[1]
	if !ff.Exist(cfgFile) {
		log.Fatal().Str("config", cfgFile).Msg("config file does not exist")
	}

	err := godotenv.Load(cfgFile)
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	runner.Start()
}

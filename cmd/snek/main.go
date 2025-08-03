package main

import (
	"flag"
	"github.com/kaiiorg/snek/pkg/game"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gopkg.in/natefinch/lumberjack.v2"
	"runtime/debug"
	"time"
)

var (
	logLevel    = flag.String("log-level", "info", "Default zerolog log level: trace, debug, info, warn, error, panic, none, etc")
	logJson     = flag.Bool("log-json", false, "Log JSON instead of pretty printing logs")
	logPath     = flag.String("log-path", "./snek.log", "Where to save log files to")
	logMaxSize  = flag.Int("log-max-size", 5, "Max size of log files in MB")
	logMaxFiles = flag.Int("log-max-files", 3, "How many file to rotate logs between")
)

func main() {
	flag.Parse()
	configureLogging()
	game.New().Run()
}

func configureLogging() {
	// Default to info level if we fail the parse the log level or were given an empty string
	level, err := zerolog.ParseLevel(*logLevel)
	if err != nil || *logLevel == "" {
		level = zerolog.InfoLevel
	}

	// If we were given "none", set global level to nolevel and immediately return
	if *logLevel == "none" {
		zerolog.SetGlobalLevel(zerolog.NoLevel)
		return
	}

	lj := &lumberjack.Logger{
		Filename:   *logPath,
		MaxSize:    *logMaxSize,
		MaxBackups: *logMaxFiles,
	}

	if *logJson {
		zerolog.TimeFieldFormat = time.RFC3339Nano
		log.Logger = log.Output(lj)
	} else {
		log.Logger = log.Output(zerolog.ConsoleWriter{
			Out:        lj,
			NoColor:    true,
			TimeFormat: time.RFC3339Nano,
		})
	}

	bi, ok := debug.ReadBuildInfo()
	if ok {
		log.Info().Str("version", bi.Main.Version).Msg("Snek")
	} else {
		log.Warn().Msg("Failed to read Snek version!")
	}

	zerolog.SetGlobalLevel(level)
}

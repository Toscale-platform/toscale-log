package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func Get() zerolog.Logger {
	return log.Output(zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "02 Jan 15:04:05",
	}).With().Caller().Logger()
}

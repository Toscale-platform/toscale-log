package toscale_log

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "02 Jan 15:04:05"}).With().Caller().Logger()
}

func m(msg interface{}) string {
	return fmt.Sprintf("%v", msg)
}

func SetLevel(l zerolog.Level) {
	zerolog.SetGlobalLevel(l)
}

func Panic(msg string) {
	log.Panic().Msg(m(msg))
}

func Fatal(msg string) {
	log.Fatal().Msg(m(msg))
}

func Error(msg string) {
	log.Error().Msg(m(msg))
}

func Warn(msg string) {
	log.Warn().Msg(m(msg))
}

func Info(msg interface{}) {
	log.Info().Msg(m(msg))
}

func Debug(msg string) {
	log.Debug().Msg(m(msg))
}

func Trace(msg string) {
	log.Trace().Msg(m(msg))
}

package log

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

func Panic(msg interface{}) {
	log.Panic().Msg(m(msg))
}

func Fatal(msg interface{}) {
	log.Fatal().Msg(m(msg))
}

func Error(msg interface{}) {
	log.Error().Msg(m(msg))
}

func Warn(msg interface{}) {
	log.Warn().Msg(m(msg))
}

func Info(msg interface{}) {
	log.Info().Msg(m(msg))
}

func Debug(msg interface{}) {
	log.Debug().Msg(m(msg))
}

func Trace(msg interface{}) {
	log.Trace().Msg(m(msg))
}

package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func GetLogger() zerolog.Logger {
	return log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "02 Jan 15:04:05"}).With().Caller().Logger()
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

func Panicf(format string, v ...interface{}) {
	log.Panic().Msgf(format, v...)
}

func Fatal(msg interface{}) {
	log.Fatal().Msg(m(msg))
}

func Fatalf(format string, v ...interface{}) {
	log.Fatal().Msgf(format, v...)
}

func Error(msg interface{}) {
	log.Error().Msg(m(msg))
}

func Errorf(format string, v ...interface{}) {
	log.Error().Msgf(format, v...)
}

func Warn(msg interface{}) {
	log.Warn().Msg(m(msg))
}

func Warnf(format string, v ...interface{}) {
	log.Warn().Msgf(format, v...)
}

func Info(msg interface{}) {
	log.Info().Msg(m(msg))
}

func Infof(format string, v ...interface{}) {
	log.Info().Msgf(format, v...)
}

func Debug(msg interface{}) {
	log.Debug().Msg(m(msg))
}

func Debugf(format string, v ...interface{}) {
	log.Debug().Msgf(format, v...)
}

func Trace(msg interface{}) {
	log.Trace().Msg(m(msg))
}

func Tracef(format string, v ...interface{}) {
	log.Trace().Msgf(format, v...)
}

package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type Logger struct {
	core zerolog.Logger
}

func GetLogger() Logger {
	return Logger{
		log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "02 Jan 15:04:05"}).With().Logger(),
	}
}

func m(msg interface{}) string {
	return fmt.Sprintf("%v", msg)
}

func (l *Logger) Panic(msg interface{}) {
	l.core.Panic().Msg(m(msg))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.core.Panic().Msgf(format, v...)
}

func (l *Logger) Fatal(msg interface{}) {
	l.core.Fatal().Msg(m(msg))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.core.Fatal().Msgf(format, v...)
}

func (l *Logger) Error(msg interface{}) {
	l.core.Error().Msg(m(msg))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.core.Error().Msgf(format, v...)
}

func (l *Logger) Warn(msg interface{}) {
	l.core.Warn().Msg(m(msg))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.core.Warn().Msgf(format, v...)
}

func (l *Logger) Info(msg interface{}) {
	l.core.Info().Msg(m(msg))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.core.Info().Msgf(format, v...)
}

func (l *Logger) Debug(msg interface{}) {
	l.core.Debug().Msg(m(msg))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.core.Debug().Msgf(format, v...)
}

func (l *Logger) Trace(msg interface{}) {
	l.core.Trace().Msg(m(msg))
}

func (l *Logger) Tracef(format string, v ...interface{}) {
	l.core.Trace().Msgf(format, v...)
}

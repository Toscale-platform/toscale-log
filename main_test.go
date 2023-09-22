package toscale_log

import (
	"errors"
	"github.com/rs/zerolog"
	"testing"
)

func TestPanic(t *testing.T) {
	Panic("Panic")
}

func TestFatal(t *testing.T) {
	Fatal("Fatal")
}

func TestError(t *testing.T) {
	Error("Error")
}

func TestWarn(t *testing.T) {
	Warn("Warn")
}

func TestInfo(t *testing.T) {
	Info("Info")
}

func TestDebug(t *testing.T) {
	Debug("Debug")
}

func TestTrace(t *testing.T) {
	Trace("Trace")
}

func TestSetLevel(t *testing.T) {
	SetLevel(zerolog.InfoLevel)
	Info("Info")
	Debug("Debug")
}

func TestConverter(t *testing.T) {
	Info("String")
	Info("String " + "String2 " + "String3")
	Info(100)
	Info([]int{1, 2, 3})
	Info([]string{"a", "b", "c"})
	Info(errors.New("new error"))
	Info(struct {
		f1 string
		f2 int
	}{
		f1: "String",
		f2: 10,
	})
}

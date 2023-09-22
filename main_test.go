package logger

import (
	"errors"
	"testing"
)

var newLogger = GetLogger()

func TestPanic(t *testing.T) {
	newLogger.Panic("Panic")
}

func TestFatal(t *testing.T) {
	newLogger.Fatal("Fatal")
}

func TestError(t *testing.T) {
	newLogger.Error("Error")
}

func TestWarn(t *testing.T) {
	newLogger.Warn("Warn")
}

func TestInfo(t *testing.T) {
	newLogger.Info("Info")
}

func TestDebug(t *testing.T) {
	newLogger.Debug("Debug")
}

func TestTrace(t *testing.T) {
	newLogger.Trace("Trace")
}

func TestConverter(t *testing.T) {
	newLogger.Info("String")
	newLogger.Info("String " + "String2 " + "String3")
	newLogger.Info(true)
	newLogger.Info(100)
	newLogger.Info([]int{1, 2, 3})
	newLogger.Info([]string{"a", "b", "c"})
	newLogger.Info(errors.New("new error"))
	newLogger.Info(struct {
		f1 string
		f2 int
	}{
		f1: "String",
		f2: 10,
	})
}

func TestFormat(t *testing.T) {
	//newLogger.Panicf("%d", 1)
	//newLogger.Fatalf("%s", "str")
	newLogger.Errorf("%f", 10.0)
	newLogger.Warnf("%s", "str")
	newLogger.Infof("%d", 10)
	newLogger.Debugf("%t", true)
	newLogger.Tracef("%p", []int{1, 2, 3})
}

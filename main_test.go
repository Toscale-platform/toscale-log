package logger

import (
	"testing"
)

var newLogger = Get()

func TestError(t *testing.T) {
	newLogger.Error().Msg("New error message")
}

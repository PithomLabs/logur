package bugsnaglog

import (
	"testing"

	"github.com/bugsnag/bugsnag-go"
	"github.com/goph/logur"
	"github.com/goph/logur/testing"
)

func TestLoggerIsBugsnagCompatible(t *testing.T) {
	_ = bugsnag.Configuration{
		Logger: New(logur.NewNoopLogger()),
	}
}

func TestHandlerLogger_Printf(t *testing.T) {
	testLogger := logur.NewTestLogger()
	logger := New(testLogger)

	logger.Printf("this is a %s", "message")

	event := logur.LogEvent{
		Level: logur.Error,
		Line:  "this is a message",
	}

	logtesting.AssertLogEvents(t, event, *(testLogger.LastEvent()))
}

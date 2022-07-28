package logging

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	filename "github.com/keepeye/logrus-filename"
	"github.com/sirupsen/logrus"
)

type LogFormatter struct {
}

func (l *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	formatter := &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime: "timestamp",
			logrus.FieldKeyMsg:  "message",
		},
	}
	return formatter.Format(entry)
}

func formatLogger(logger *logrus.Logger) *logrus.Logger {
	logger.SetFormatter(&LogFormatter{})
	filenameHook := filename.NewHook()
	filenameHook.Field = "line"
	logger.AddHook(filenameHook)
	logger.SetLevel(logrus.TraceLevel)
	return logger
}

// NewLogger instantiates a logrus logger with basic format.
func NewLogger(requestID string, actorCrn string) *logrus.Entry {
	logger := formatLogger(logrus.New())
	entry := logrus.NewEntry(logger).WithFields(logrus.Fields{
		"TenantID":  "1234-5678-9101",
		"RequestID": "1234-5678-9111",
	})
	return entry
}

type Logger interface {
	Error(format string, args ...interface{})
}
type AppLogger struct {
	logger *logrus.Entry
}

func (a AppLogger) Error(format string, args ...interface{}) {
	a.logger.Errorf(format, args...)
	go a.log(format, args...)

}

func NewAppLogger() Logger {
	return &AppLogger{
		logger: NewLogger(uuid.NewString(), uuid.NewString()),
	}
}

func (a AppLogger) log(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: "https://f3a073599a794b0d87eba6c27144fb31@o1338048.ingest.sentry.io/6608470",
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: "",
		Release:     "my-project-gaurav-dummy@1.0.0",
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Flush buffered events before the program terminates.
	// Set the timeout to the maximum duration the program can afford to wait.
	defer sentry.Flush(2 * time.Second)

	sentry.CaptureMessage(msg)
	fmt.Println("printed message to sentry")
}

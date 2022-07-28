package logging

import (
	"log"

	"github.com/getsentry/sentry-go"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type SentryConfig struct {
}

type FormatSettings struct {
	jsonFormat  bool
	bsonFormat  bool
	kafkaFormat bool
}

// CELogger is CloudEra logging wrapper
type CELogger struct {
	Logger         *logrus.Entry
	FormatSettings *FormatSettings
}

// NewCELogger ...
func NewCELogger() *CELogger {
	logger := NewLogger(uuid.NewString(), uuid.NewString())
	return &CELogger{Logger: logger}
}

func NewCELoggerWithSenity(dsn, env, release string, debug bool) *CELogger {
	logger := NewLogger(uuid.NewString(), uuid.NewString())
	err := sentry.Init(sentry.ClientOptions{
		// Either set your DSN here or set the SENTRY_DSN environment variable.
		Dsn: dsn,
		// Either set environment and release here or set the SENTRY_ENVIRONMENT
		// and SENTRY_RELEASE environment variables.
		Environment: env,
		Release:     release,
		// Enable printing of SDK debug messages.
		// Useful when getting started or trying to figure something out.
		Debug: debug,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
	})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	return &CELogger{Logger: logger}
}

func NewCELoggerWithFomatter(jsonFormat, bsonFormat, kafka bool) *CELogger {
	logger := NewLogger(uuid.NewString(), uuid.NewString())
	formatSettings := FormatSettings{
		jsonFormat:  jsonFormat,
		bsonFormat:  bsonFormat,
		kafkaFormat: kafka,
	}
	return &CELogger{Logger: logger, FormatSettings: &formatSettings}
}

package logger

import (
	"os"

	"github.com/getsentry/sentry-go"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type SentryHook struct{}

func NewSentryHook() (zerolog.Hook, error) {
	if err := sentry.Init(sentry.ClientOptions{
		Environment: "development",
		Dsn:         "",
	}); err != nil {
		_ = err
		return nil, err
	}

	return &SentryHook{}, nil
}

func (h SentryHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
	if level == zerolog.NoLevel || level == zerolog.Disabled {
		return
	}

	if level > zerolog.WarnLevel {
		se := sentry.NewEvent()
		se.Tags["level"] = zerolog.LevelFieldMarshalFunc(level)
		se.Message = msg
		sentry.CaptureEvent(se)
	}
}

type Config struct {
	Type string
}

type Logger struct{}

func New(cfg *Config) *Logger {
	if cfg.Type != "json" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
	zerolog.TimestampFieldName = "timestamp"
	return &Logger{}
}

func (l *Logger) AddHook(hooks ...zerolog.Hook) {
	for _, hook := range hooks {
		log.Logger = log.Logger.Hook(hook)
	}
}

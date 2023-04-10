package configs

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// Hook is a hook that writes logs of specified LogLevels to specified Writer
type Hook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
	Formatter logrus.Formatter
}

type FieldMap struct {
	FieldKeyTime  string
	FieldKeyLevel string
	FieldKeyMsg   string
	FieldKeyFunc  string
}

// Fire will be called when some logging function is called with current hook
// It will format log entry to string and write it to appropriate writer
func (h *Hook) Fire(entry *logrus.Entry) error {
	line, err := h.Formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.Writer.Write(line)
	return err
}

// Levels define on which log levels this hook would trigger
func (h *Hook) Levels() []logrus.Level {
	return h.LogLevels
}

func Init() (*os.File, *logrus.Logger) {
	log := logrus.New()
	file, err := os.OpenFile("logs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Failed to create logfile" + "log.txt")
		panic(err)
	}
	log.SetOutput(io.Discard)
	log.SetReportCaller(true)

	log.AddHook(&Hook{
		Writer: os.Stderr,
		LogLevels: []logrus.Level{
			logrus.DebugLevel,
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
		},
		Formatter: &logrus.TextFormatter{
			FullTimestamp:   true,
			ForceColors:     true,
			TimestampFormat: "2006-01-02 15:04:05",
			DisableColors:   false,
		},
	})

	log.AddHook(&Hook{
		Writer: file,
		LogLevels: []logrus.Level{
			logrus.DebugLevel,
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
		},
		Formatter: &logrus.JSONFormatter{},
	})

	return file, log
}

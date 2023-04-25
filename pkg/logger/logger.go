package logger

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

type WriterHook struct {
	Writer    io.Writer
	LogLevels []logrus.Level
}

func init() {
	os.Setenv("TZ", "Europe/Moscow")
}
func (hook *WriterHook) Fire(entry *logrus.Entry) error {
	line, err := entry.String()
	if err != nil {
		return err
	}
	_, err = hook.Writer.Write([]byte(line))
	return err
}

func (hook *WriterHook) Levels() []logrus.Level {
	return hook.LogLevels
}

func New() *logrus.Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "02.01.2006 15:04:05",
	})

	if err := os.MkdirAll("tmp", 0755); err != nil && !os.IsExist(err) {
		log.Fatalf("can't creaete tmp dir: %v", err)
	}

	f, err := os.OpenFile("tmp/service.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	logger.SetOutput(io.Discard)
	logger.AddHook(&WriterHook{
		Writer: f,
		LogLevels: []logrus.Level{
			logrus.PanicLevel,
			logrus.FatalLevel,
			logrus.ErrorLevel,
			logrus.WarnLevel,
			logrus.InfoLevel,
			logrus.TraceLevel,
		},
	})
	logger.AddHook(&WriterHook{
		Writer: os.Stdout,
		LogLevels: []logrus.Level{
			logrus.InfoLevel,
			logrus.DebugLevel,
			logrus.ErrorLevel,
		},
	})

	return logger
}

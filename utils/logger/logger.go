package logger

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		PadLevelText:  true,
	})
	Logger.SetOutput(os.Stdout)
	Logger.SetLevel(logrus.DebugLevel)
}

func Info(msg string) {
	Logger.Info(msg)
}

func Warn(msg string) {
	Logger.Warn(msg)
}

func Error(msg string, err error) {
	if err != nil {
		Logger.Error(fmt.Sprintf("%s | %v", msg, err))
	} else {
		Logger.Error(msg)
	}
}

func Fatal(msg string, err error) {
	if err != nil {
		Logger.Fatal(fmt.Sprintf("%s | %v", msg, err))
	} else {
		Logger.Fatal(msg)
	}
}

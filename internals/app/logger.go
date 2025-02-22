package app

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

func init() {
	Logger = logrus.New()

	Logger.WithField("app", "lb-mini")
	Logger.SetLevel(LOG_LEVEL)
	Logger.SetOutput(os.Stdout)
}

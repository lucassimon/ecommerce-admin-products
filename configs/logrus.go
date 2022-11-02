package configs

import (
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyTime:  "timestamp",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyFunc:  "caller",
			logrus.FieldKeyMsg:   "message",
		},
	})

	if os.Getenv("ENVIRONMENT") == "live" {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.DebugLevel)
	}
}

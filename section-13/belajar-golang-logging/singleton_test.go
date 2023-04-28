package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestSingleton(t *testing.T) {
	logrus.Info("Hello Logrus")
	logrus.Warn("Hello Logrus")

	logrus.SetFormatter(new(logrus.JSONFormatter))

	logrus.Info("Hello Logrus")
	logrus.Warn("Hello Logrus")
}

package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")
}

package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	entry := logrus.NewEntry(logger)
	entry.WithField("username", "ymanshur")
	entry.Info("Hello Entry")
}

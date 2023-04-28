package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"testing"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	logger.
		WithField("username", "ymanshur").
		WithField("name", "Yusuf Manshur").
		Info("Hello World")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))

	logger.WithFields(logrus.Fields{
		"username": "ymanshur",
		"name":     "Yusuf Manshur",
	}).Info("Hello World")
}

package belajar_golang_logging

import (
	"github.com/sirupsen/logrus"
	"os"
	"testing"
)

func TestOutput(t *testing.T) {
	logger := logrus.New()

	// set writer
	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_RDONLY, 0666)
	logger.SetOutput(file)

	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")
}

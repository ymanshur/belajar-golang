package belajar_golang_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestLogger(t *testing.T) {
	logger := logrus.New()

	logger.Println("Hello Logger")
	fmt.Println("Hello Logger")
}

func TestLever(t *testing.T) {
	logger := logrus.New()

	logger.Trace("This is Trace") // not shown
	logger.Debug("This is Debug") // not shown

	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")

	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
	logger.Fatal("This is Fatal")
	logger.Panic("This is Panic")
}

package belajar_golang_logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"testing"
)

type SampleHook struct {
}

func (s *SampleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.WarnLevel}
}

func (s *SampleHook) Fire(entry *logrus.Entry) error {
	// shown first before the log
	fmt.Println("Sample Hook", entry.Level, entry.Message)
	return nil
}

func TestHook(t *testing.T) {
	logger := logrus.New()
	logger.AddHook(new(SampleHook))

	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}

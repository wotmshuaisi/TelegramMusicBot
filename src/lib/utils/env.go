package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

// GetEnvWithDefault ...
func GetEnvWithDefault(name, d string) string {
	s := os.Getenv(name)
	if s == "" {
		return d
	}
	return s
}

// GetEnvWithFatal ...
func GetEnvWithFatal(name string) string {
	s := os.Getenv(name)
	if s == "" {
		logrus.Fatalf("environment " + name + " does not exist")
	}
	return s
}

package utils

import (
	"os"
	"testing"
)

func TestGetEnvWithDefault(t *testing.T) {
	s := GetEnvWithDefault("A", "A")
	if s != "A" {
		t.Fatal(s)
	}
	os.Setenv("A", "s")
	defer os.Unsetenv("A")
	s = GetEnvWithDefault("A", "s")
	if s != "s" {
		t.Fatal(s)
	}
}

func TestGetEnvWithFatal(t *testing.T) {
	os.Setenv("A", "s")
	defer os.Unsetenv("A")
	s := GetEnvWithFatal("A")
	if s != "s" {
		t.Fatal(s)
	}
}

package utils

import (
	"testing"
)

func TestGetJSON(t *testing.T) {
	b, err := HTTPGetJSON("https://www.baidu.com")
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal()
	}
}

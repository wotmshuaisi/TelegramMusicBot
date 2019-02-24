package utils

import (
	"testing"
)

func TestGetJSON(t *testing.T) {
	return
	b, err := HTTPGetJSON("https://api.imjad.cn/cloudmusic/?type=song&id=35847388&search_type=1")
	if err != nil {
		t.Fatal(err)
	}
	if len(b) == 0 {
		t.Fatal()
	}
}

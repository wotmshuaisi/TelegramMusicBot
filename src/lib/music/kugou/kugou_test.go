package kuwo

import (
	"testing"
)

func Test_handler_List(t *testing.T) {
	h := NewAPI()
	res, err := h.List("hello")
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Fatal(res)
	}
}

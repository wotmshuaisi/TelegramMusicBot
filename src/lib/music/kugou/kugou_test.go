package kugou

import (
	"testing"
)

func Test_handler_List(t *testing.T) {
	h := NewAPI()
	res, err := h.ListItem("hello")
	if err != nil {
		t.Fatal(err)
	}
	if res == nil {
		t.Fatal(res)
	}

	for _, v := range *res {
		if v.URL == "" || v.URL == "ERROR" || v.Title == "" || v.ID == "" {
			t.Fatal(v)
		}
	}

}

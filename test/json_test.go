package test

import (
	"encoding/json"
	"testing"
)

type Message struct {
	Name     string `json:"name_json"`
	Body     string
	Time     int64
	Case     string `json:"case_json,omitempty"`
	Password string `json:"-"`
}

func TestJson(t *testing.T) {
	m := Message{"Alice", "Hello", 1294706395881547000, "", "password_string"}
	b, err := json.Marshal(m)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(string(b))
	}

	var n Message
	errUnmarshal := json.Unmarshal(b, &n)
	if errUnmarshal != nil {
		t.Log(errUnmarshal)
	} else {
		t.Log(n)
	}
}

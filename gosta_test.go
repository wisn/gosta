package gosta

import (
	"testing"
)

func TestNew(t *testing.T) {
	cfg := Config{
		Host: "http://shroudoftheavatar.com",
		Port: 9200,
	}

	_, err := New(cfg)
	if err != nil {
		t.Error(err)
	}
}

package fitters

import (
	"testing"
	"time"
)

func TestWindowCreation(t *testing.T) {
	w, err := NewWindow(time.Duration(60)*time.Second, time.Duration(1)*time.Second)
	if err != nil {
		t.Fatalf("NewWindow didn't create a window", err)
	}

	if w.Buckets() != 60 {
		t.Fatalf("NewWindow didn't create a 60 bucket window")
	}
}

func TestWindowFailure(t *testing.T) {
	w, err := NewWindow(time.Duration(1)*time.Second, time.Duration(60)*time.Second)
	if w != nil || err == nil {
		t.Fatalf("NewWindow created an invalid window")
	}
}

package main

import "testing"

func TestGreet(t *testing.T) {
	msg := greet()
	if len(msg) > 0 {
		t.Fatalf("invalid content")
	}
}

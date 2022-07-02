package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	in, out := makeUI()
	if out.Text != "" {
		t.Error("Incorrect initial greeting")
	}

	//mock user input
	test.Type(in, "Andy")
	if out.Text != "Greeting: Hello Andy!" {
		t.Error("Incorrect user greeting")
	}

}

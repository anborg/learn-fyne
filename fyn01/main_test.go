package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
)

func TestGreeting(t *testing.T) {
	username, _, greeting, _ := makeUI()
	if greeting.Text != "" {
		t.Error("Incorrect initial greeting")
	}

	//mock user input
	test.Type(username, "Andy")
	if greeting.Text != "Greeting: Hello Andy!" {
		t.Error("Incorrect user greeting")
	}

}

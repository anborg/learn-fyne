package main

import (
	"testing"

	"fyne.io/fyne/v2/test"
	"fyne.io/fyne/v2/widget"
)

func TestGreeting(t *testing.T) {
	widgetMap := makeUI()
	username := widgetMap["username"].(*widget.Entry)
	greeting := widgetMap["greeting"].(*widget.Label)
	if greeting.Text != "" {
		t.Error("Incorrect initial greeting")
	}

	//mock user input
	test.Type(username, "Andy")
	if greeting.Text != "Greeting: Hello Andy!" {
		t.Error("Incorrect user greeting")
	}

}

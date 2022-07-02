package main

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Clock")
	clock := widget.NewLabel("Time")
	updateTime(clock)
	w.SetContent(container.NewVBox(
		makeUI()))
	w.SetMaster()
	//To uodate time constantly, create a goroutine
	go func() { //Place this code before  Run()
		for range time.Tick(time.Second) {
			updateTime(clock)
		}
	}()

	w.Show()
	w2 := a.NewWindow("target")
	w2.SetContent(widget.NewButton("Open New", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third Window"))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	//make one window master

	a.Run() //Event-loop, or Run-loop
	tidyUp()
}
func makeUI() (in *widget.Entry, out *widget.Label) {
	in = widget.NewEntry()
	out = widget.NewLabel("")

	in.OnChanged = func(content string) {
		out.SetText("Greeting: Hello " + content + "!")
	}
	return
}

func updateTime(clock *widget.Label) {
	formatted := time.Now().Format("Time: 03:04:05")
	clock.SetText(formatted)
}
func tidyUp() {
	fmt.Println("Exited")
}

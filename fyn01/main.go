package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.NewWithID("com.example.tutorial.preferences")
	a.Preferences().SetBool("Boolean", true)

	//to demo preferences
	var timeout time.Duration
	timeoutSelector := widget.NewSelect([]string{"10", "20", "30", "120"},
		func(selected string) {
			i, _ := strconv.Atoi(selected)
			timeout = time.Duration(i) * time.Second
			a.Preferences().SetString("AppTimeout", selected)
		},
	) //timeoutselcector
	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "10"))

	w := a.NewWindow("Timeout")
	clock := widget.NewLabel("Time")
	updateTime(clock)

	//Destop or mobile
	if a.Driver().Device().IsMobile() {
		mycontainer := phoneLayout(makeUI())
		mycontainer.Add(timeoutSelector)
		mycontainer.Add(clock)
		w.SetContent(mycontainer)
	} else {

		w.SetContent(desktopLayout(makeUI()))
	}

	w.SetMaster() // make one window as master

	w.Show()
	w2 := a.NewWindow("target")
	w2.SetContent(widget.NewButton("Open New", func() {
		w3 := a.NewWindow("Third")
		w3.SetContent(widget.NewLabel("Third Window"))
		w3.Show()
	}))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	//Backgroud processes
	//Use goroutines for background tasks
	go func() { //Place this code before  Run()
		//Update time
		for range time.Tick(time.Second) {
			updateTime(clock)
		}

	}()

	go func() {
		//Quit application based on preference timeset
		fmt.Println("Time to sleep", timeout.String())
		time.Sleep(timeout)
		fmt.Println("Time to close the app?!")
		a.Quit()
	}()

	a.Run() //Event-loop, or Run-loop
	tidyUp()
}

func phoneLayout(widgetMap map[string]fyne.CanvasObject) *fyne.Container {
	return container.NewVBox(
		widgetMap["username"],
		widgetMap["password"],
		widgetMap["button"],
		layout.NewSpacer(),
		widgetMap["greeting"],
	)
}

//username *widget.Entry, password *widget.Entry, greeting *widget.Label, button *widget.Button
func desktopLayout(widgetMap map[string]fyne.CanvasObject) *fyne.Container {
	return container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3, //second row spint into 3 col
			layout.NewSpacer(),
			container.NewVBox(
				widgetMap["username"],
				widgetMap["password"],
				widgetMap["button"],
				layout.NewSpacer(),
				widgetMap["greeting"],
			),
		),
		layout.NewSpacer(),
	)
}

//sername *widget.Entry, password *widget.Entry, greeting *widget.Label, button *widget.Button,
func makeUI() (widgetMap map[string]fyne.CanvasObject) {

	username := widget.Entry{PlaceHolder: "Username"}
	password := widget.Entry{PlaceHolder: "Password", Password: true}
	greeting := widget.NewLabel("")
	button := &widget.Button{Text: "Login", Icon: theme.ConfirmIcon()}

	username.OnChanged = func(content string) {
		greeting.SetText("Greeting: Hello " + content + "!")
	}

	widgetMap = map[string]fyne.CanvasObject{
		"username": username,
		"password": password,
		"greeting": greeting,
		"button":   button,
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

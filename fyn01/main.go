package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.NewWithID("com.example.tutorial.preferences")
	a.Preferences().SetBool("Boolean", true)

	//to demo preferences
	var timeout time.Duration
	timeoutSelector := widget.NewSelect([]string{"10", "20", "60", "120"},
		func(selected string) {
			i, _ := strconv.Atoi(selected)
			timeout = time.Duration(i) * time.Second
			a.Preferences().SetString("AppTimeout", selected)
		},
	) //timeoutselcector
	timeoutSelector.SetSelected(a.Preferences().StringWithFallback("AppTimeout", "60"))

	w := a.NewWindow("Timeout")
	//demo databinding
	timeStr := binding.NewString()
	clock := widget.NewLabelWithData(timeStr)

	var mycontainer *fyne.Container
	//Destop or mobile
	var mywidgetMap map[string]interface{} = makeUI()
	mywidgetMap["timeoutSelector"] = timeoutSelector
	mywidgetMap["clock"] = clock
	if a.Driver().Device().IsMobile() {
		mycontainer = phoneLayout(mywidgetMap)
	} else {
		mycontainer = desktopLayout(mywidgetMap)
	}
	w.SetContent(mycontainer)

	w.SetMaster() // make one window as master
	w.Show()

	//to demo systray --
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Fyn01",
			fyne.NewMenuItem("Show", func() {
				log.Println("Systray::Show clicked - window will be shown")
				w.Show()
			})) //menu
		desk.SetSystemTrayMenu(m)
		//also intercept and hide, instead of quit

		w.SetCloseIntercept(func() {
			log.Println("MainWindow::close icon clicked - app will be hidden")
			w.Hide()
		})

	} //desk

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
			formatted := time.Now().Format("Time: 03:04:05")
			log.Println("In timer FOR loop,", formatted)
			timeStr.Set(formatted)
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

func phoneLayout(widgetMap map[string]interface{}) *fyne.Container {
	return container.NewVBox(
		widgetMap["username"].(*widget.Entry),
		widgetMap["password"].(*widget.Entry),
		widgetMap["button"].(fyne.CanvasObject),
		layout.NewSpacer(),
		widgetMap["greeting"].(fyne.CanvasObject),
		widgetMap["timeoutSelector"].(fyne.CanvasObject),
		widgetMap["clock"].(fyne.CanvasObject),
	)
}

//username *widget.Entry, password *widget.Entry, greeting *widget.Label, button *widget.Button
func desktopLayout(widgetMap map[string]interface{}) *fyne.Container {
	return container.NewGridWithRows(3,
		layout.NewSpacer(),
		container.NewGridWithColumns(3, //second row spint into 3 col
			layout.NewSpacer(),
			container.NewVBox(
				widgetMap["username"].(*widget.Entry), //how can I avoid casting?
				widgetMap["password"].(*widget.Entry),
				widgetMap["button"].(fyne.CanvasObject),
				layout.NewSpacer(),
				widgetMap["greeting"].(fyne.CanvasObject),
				widgetMap["timeoutSelector"].(fyne.CanvasObject),
				widgetMap["clock"].(fyne.CanvasObject),
			),
		),
		layout.NewSpacer(),
	)
}

//sername *widget.Entry, password *widget.Entry, greeting *widget.Label, button *widget.Button,
func makeUI() (widgetMap map[string]interface{}) {
	strGreet := binding.NewString()
	strGreet.Set("Hi!")

	username := widget.NewEntryWithData(strGreet)
	username.SetPlaceHolder("Username")
	password := &widget.Entry{PlaceHolder: "Password", Password: true}
	greeting := widget.NewLabelWithData(strGreet)
	button := &widget.Button{Text: "Login", Icon: theme.ConfirmIcon()}

	// username.OnChanged = func(content string) {
	// 	greeting.SetText("Greeting: Hello " + content + "!")
	// }

	widgetMap = map[string]interface{}{
		"username": username,
		"password": password,
		"greeting": greeting,
		"button":   button,
	}
	return
}

func tidyUp() {
	fmt.Println("Exited")
}

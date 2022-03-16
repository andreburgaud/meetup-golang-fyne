package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateTime(clock *widget.Label) {
	t := time.Now().Format("03:04:05")
	clock.SetText(t)
}

func main() {
    a := app.New()
    w := a.NewWindow("Clock")

    clock := widget.NewLabel("")
    updateTime(clock)

    w.SetContent(clock)
    go func() {
    	for range time.Tick(time.Second) {
    		updateTime(clock)
    	}
    }()

    w.ShowAndRun()
}

package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello GoMN 😊")
	w.SetContent(widget.NewLabel("I hope you are all Fyne"))
	w.ShowAndRun()
}

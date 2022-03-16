package main

import (
	"fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
)

func nameInputUI() (*widget.Label, *widget.Entry) {
	out := widget.NewLabel("Anonymous")
	in := widget.NewEntry()

	in.OnChanged = func(content string) {
		if content == "" {
			content = "Anonymous"
		}
		out.SetText(content)
	}
	return out, in
}

func main() {
    a := app.New()
    w := a.NewWindow("Hello GoMN 😊")

    w.SetContent(container.NewVBox(nameInputUI()))
    w.Resize(fyne.NewSize(200, 100))
    w.ShowAndRun()
}

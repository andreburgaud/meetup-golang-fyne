package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/widget"
)

func main() {
    a := app.New()
    
    w1 := a.NewWindow("One")
    w1.SetContent(widget.NewLabel("Window One"))
    w1.SetMaster()
    w1.Show()

    w2 := a.NewWindow("Two")
    w2.SetContent(widget.NewButton("Open Window", func() {
        w3 := a.NewWindow("Three")
        w3.SetContent(widget.NewLabel("Window Three"))
        w3.Show()
    }))

    w2.Resize(fyne.NewSize(100, 100))
    w2.Show()

    a.Run()
}

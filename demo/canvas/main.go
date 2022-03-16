// From Fyne document example: https://developer.fyne.io/explore/canvas.html
package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Canvas")
	c := w.Canvas()

	blue := color.NRGBA{R: 0, G: 87, B: 184, A: 255}
	yellow := color.NRGBA{R: 254, G: 221, B: 0, A: 255}
	rect := canvas.NewRectangle(blue)
	c.SetContent(rect)

	go func() {
		for {
			time.Sleep(time.Second)
			rect.FillColor = yellow
			rect.Refresh()
			time.Sleep(time.Second)
			rect.FillColor = blue
			rect.Refresh()
		}
	}()

	w.Resize(fyne.NewSize(400, 150))
	w.ShowAndRun()
}

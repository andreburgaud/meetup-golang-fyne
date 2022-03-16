// From Fyne document example: https://developer.fyne.io/explore/canvas.html
package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func getBlueAlpha(alpha uint8) color.NRGBA {
	return color.NRGBA{R: 0, G: 87, B: 184, A: alpha}
}

func getYellowAlpha(alpha uint8) color.NRGBA {
	return color.NRGBA{R: 254, G: 221, B: 0, A: alpha}
}

func main() {
	a := app.New()
	w := a.NewWindow("Flag")
	c := w.Canvas()

	rectUp := canvas.NewRectangle(color.White)
	rectUp.SetMinSize(fyne.NewSize(400, 80))
	rectUp.Refresh()
	rectDn := canvas.NewRectangle(color.White)
	rectDn.SetMinSize(fyne.NewSize(400, 80))
	rectDn.Refresh()
	c.SetContent(container.New(layout.NewVBoxLayout(), rectUp, rectDn))

	go func() {
		var n uint8 = 1
		for n < 255 {
			time.Sleep(30 * time.Millisecond)
			rectUp.FillColor = getBlueAlpha(n)
			rectUp.Refresh()
			rectDn.FillColor = getYellowAlpha(n)
			rectDn.Refresh()
			n++
		}
	}()

	w.SetFixedSize(true)
	w.ShowAndRun()
}

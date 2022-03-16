// From https://github.com/fyne-io/examples/tree/develop/clock
// Released under the BSD 3-Clause license

package main

import (
	"image/color"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type clockLayout struct {
	hour, minute, second     *canvas.Line
	hourDot, secondDot, face *canvas.Circle
	canvas                   fyne.CanvasObject
	stop                     bool
}

func (c *clockLayout) rotate(hand *canvas.Line, middle fyne.Position, facePosition float64, offset, length float32) {
	rotation := math.Pi * 2 / 60 * facePosition
	x2 := length * float32(math.Sin(rotation))
	y2 := -length * float32(math.Cos(rotation))

	offX := float32(0)
	offY := float32(0)
	if offset > 0 {
		offX += offset * float32(math.Sin(rotation))
		offY += -offset * float32(math.Cos(rotation))
	}

	hand.Position1 = fyne.NewPos(middle.X+offX, middle.Y+offY)
	hand.Position2 = fyne.NewPos(middle.X+offX+x2, middle.Y+offY+y2)
	hand.Refresh()
}

func (c *clockLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	dotRadius := radius / 12
	smallDotRadius := dotRadius / 8

	stroke := diameter / 40
	midStroke := diameter / 90
	smallStroke := diameter / 200

	size = fyne.NewSize(diameter, diameter)
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	topleft := fyne.NewPos(middle.X-radius, middle.Y-radius)

	c.face.Resize(size)
	c.face.Move(topleft)

	c.hour.StrokeWidth = stroke
	c.rotate(c.hour, middle, float64((time.Now().Hour()%12)*5)+(float64(time.Now().Minute())/12), dotRadius, radius/2)
	c.minute.StrokeWidth = midStroke
	c.rotate(c.minute, middle, float64(time.Now().Minute())+(float64(time.Now().Second())/60), dotRadius, radius*.9)
	c.second.StrokeWidth = smallStroke
	c.rotate(c.second, middle, float64(time.Now().Second()), 0, radius-3)

	c.hourDot.StrokeWidth = stroke
	c.hourDot.Resize(fyne.NewSize(dotRadius*2, dotRadius*2))
	c.hourDot.Move(fyne.NewPos(middle.X-dotRadius, middle.Y-dotRadius))
	c.secondDot.StrokeWidth = smallStroke
	c.secondDot.Resize(fyne.NewSize(smallDotRadius*2, smallDotRadius*2))
	c.secondDot.Move(fyne.NewPos(middle.X-smallDotRadius, middle.Y-smallDotRadius))
	c.face.StrokeWidth = smallStroke
}

func (c *clockLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 200)
}

func (c *clockLayout) render() *fyne.Container {
	red := color.RGBA{R: 254, G: 0, B: 0, A: 255}
	blue := color.NRGBA{R: 0, G: 87, B: 184, A: 255}
	yellow := color.NRGBA{R: 254, G: 221, B: 0, A: 255}

	c.hourDot = &canvas.Circle{StrokeColor: yellow, StrokeWidth: 5}
	c.secondDot = &canvas.Circle{StrokeColor: red, StrokeWidth: 3}
	c.face = &canvas.Circle{StrokeColor: blue, StrokeWidth: 1}

	c.hour = &canvas.Line{StrokeColor: yellow, StrokeWidth: 5}
	c.minute = &canvas.Line{StrokeColor: blue, StrokeWidth: 3}
	c.second = &canvas.Line{StrokeColor: red, StrokeWidth: 1}

	container := container.NewWithoutLayout(c.hourDot, c.secondDot, c.face, c.hour, c.minute, c.second)
	container.Layout = c

	c.canvas = container
	return container
}

func (c *clockLayout) animate(co fyne.CanvasObject) {
	tick := time.NewTicker(time.Second)
	go func() {
		for !c.stop {
			c.Layout(nil, co.Size())
			canvas.Refresh(c.canvas)
			<-tick.C
		}
	}()
}

func show(win fyne.Window) fyne.CanvasObject {
	clock := &clockLayout{}
	content := clock.render()
	go clock.animate(content)
	return content
}

func main() {
	a := app.New()
	w := a.NewWindow("Analog Clock")
	c := container.NewMax()
	w.SetContent(c)
	c.Objects = []fyne.CanvasObject{show(w)}
	w.ShowAndRun()
}

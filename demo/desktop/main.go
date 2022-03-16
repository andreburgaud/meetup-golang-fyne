package main

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

const (
	appName    = "Fyne Demo App"
	appVersion = "1.0"
)

var themeColor string

// CUSTOM THEME
type customTheme struct{}

func (m customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {

		if themeColor == "Default" {

			return theme.DefaultTheme().Color(name, variant)
		}

		if themeColor == "Blue" {
			return color.NRGBA{R: 0, G: 87, B: 184, A: 255}
		}

		if themeColor == "Yellow" {
			return color.NRGBA{R: 254, G: 221, B: 0, A: 255}
		}
	}
	return theme.DefaultTheme().Color(name, variant)
}

func (m customTheme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(style)
}

func (m customTheme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}

func (m customTheme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

// END CUSTOM THEME

// TOOLBAR
func buildToolbar(w fyne.Window) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			dialog.ShowInformation("About", fmt.Sprintf("%s version %s", appName, appVersion), w)
		}),
	)
	return toolbar
}

// SELECT (DROPDOWN)
func buildColorSelect(color string) *widget.Select {
	s := widget.NewSelect([]string{"Default", "Blue", "Yellow"}, func(color string) {
		themeColor = color
	})
	s.SetSelected(color)
	return s
}

// STATUS BAR
func buildStatus(color string) *fyne.Container {
	status := container.NewHBox(layout.NewSpacer(),
		buildColorSelect(color),
		widget.NewLabel("Status:"), widget.NewLabel("Ok"),
	)
	return status
}

// FORM (email)
func buildEmailForm() *widget.Form {
	to := widget.NewEntry()
	subject := widget.NewEntry()
	message := widget.NewMultiLineEntry()

	formTo := widget.NewFormItem("Subject:", to)
	formSubject := widget.NewFormItem("Subject:", subject)
	formMessage := widget.NewFormItem("Message:", message)

	form := &widget.Form{
		Items: []*widget.FormItem{formTo, formSubject, formMessage},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Message submitted to:", to.Text)
			log.Println("Message Subject     :", subject.Text)
			log.Println("Message Text        :", message.Text)
		},
		SubmitText: "Send",
	}
	return form
}

// FORM (note)
func buildNoteForm() *widget.Form {
	note := widget.NewMultiLineEntry()
	noteItem := widget.NewFormItem("Note:", note)

	form := &widget.Form{
		Items: []*widget.FormItem{noteItem},
		OnSubmit: func() {
			if len(note.Text) > 0 {
				log.Println(note.Text)
			}
		},
		SubmitText: "Save",
	}
	return form
}

// TABS
func buildTabs() *container.AppTabs {
	emailForm := buildEmailForm()
	noteForm := buildNoteForm()
	tabs := container.NewAppTabs(
		container.NewTabItem("Email", emailForm),
		container.NewTabItem("Note", noteForm),
	)
	return tabs
}

// ENTRYPOINT
func main() {
	themeColor = "Default"
	a := app.New()
	var _ fyne.Theme = (*customTheme)(nil) // Assert interface fully implemented
	a.Settings().SetTheme(&customTheme{})

	w := a.NewWindow(appName)
	w.SetMainMenu(fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Quit", func() {
				w.Close()
			}),
		),
	))
	toolbar := buildToolbar(w)
	tabs := buildTabs()
	status := buildStatus(themeColor)
	content := container.New(layout.NewBorderLayout(toolbar, status, nil, nil),
		toolbar, status, tabs)
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

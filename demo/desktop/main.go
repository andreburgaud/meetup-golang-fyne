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

// Custom Theme
type customTheme struct{}

func (m customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	if name == theme.ColorNameBackground {
		cfg, err := GetAppConfig()

		if err != nil {
			displayError(err)
			return theme.DefaultTheme().Color(name, variant)
		}

		if cfg.Color == "Default" {
			return theme.DefaultTheme().Color(name, variant)
		}

		if cfg.Color == "Blue" {
			return color.NRGBA{R: 0, G: 87, B: 184, A: 255}
		}

		if cfg.Color == "Yellow" {
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

// End of Custom Theme

func buildToolbar(w fyne.Window) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			dialog.ShowInformation("About", fmt.Sprintf("%s version %s", appName, appVersion), w)
		}),
	)
	return toolbar
}

func displayError(err error) {
	fmt.Println(err)
}

func onColorChange(color string) {
	err := SaveColor(color)
	if err != nil {
		displayError(err)
		return
	}

	cfg, err := GetAppConfig()
	if err != nil {
		displayError(err)
		return
	}

	cfg.Print()
}

func buildColorSelect(color string) *widget.Select {
	s := widget.NewSelect([]string{"Default", "Blue", "Yellow"}, onColorChange)
	s.SetSelected(color)
	return s
}

func buildStatus(color string) *fyne.Container {
	status := container.NewHBox(layout.NewSpacer(),
		buildColorSelect(color),
		widget.NewLabel("Status:"), widget.NewLabel("Ok"),
	)
	return status
}

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

func buildTabs() *container.AppTabs {
	emailForm := buildEmailForm()
	noteForm := buildNoteForm()
	tabs := container.NewAppTabs(
		container.NewTabItem("Email", emailForm),
		container.NewTabItem("Note", noteForm),
	)
	return tabs
}

func main() {
	cfg, err := GetAppConfig()
	if err != nil {
		fmt.Println(err) // Display error in dialog bix
	}
	cfg.Print()

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
	status := buildStatus(cfg.Color)
	content := container.New(layout.NewBorderLayout(toolbar, status, nil, nil),
		toolbar, status, tabs)
	w.SetContent(content)
	w.Resize(fyne.NewSize(800, 600))
	w.ShowAndRun()
}

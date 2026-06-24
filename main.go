package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Femto Torrent")

	magnetInput := widget.NewEntry()
	magnetInput.SetPlaceHolder("Paste magnet link...")

	addButton := widget.NewButton("Add", func() {
		// add torrent
	})

	importButton := widget.NewButton("Import from file", func() {
		filter := storage.NewExtensionFileFilter([]string{".torrent", ".magnet"})

		dialog.ShowFileOpen(
			func(uri fyne.URIReadCloser, err error) {
				if err != nil || uri == nil {
					return
				}

				println("Selected:", uri.URI().String())
			},
			w,
		)

		_ = filter
	})

	buttonRow := container.NewHBox(
		addButton,
		importButton,
	)

	top := container.NewBorder(
		nil,
		nil,
		nil,
		buttonRow,
		magnetInput,
	)

	torrentList := container.NewVScroll(
		container.NewVBox(
			widget.NewLabel("Ubuntu ISO"),
			widget.NewProgressBar(),
		),
	)

	status := widget.NewLabel("Idling...")

	content := container.NewBorder(
		top,
		status,
		nil,
		nil,
		torrentList,
	)

	w.SetContent(content)

	w.SetFixedSize(true)

	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}

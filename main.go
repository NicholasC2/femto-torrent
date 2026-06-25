package main

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"femto-torrent/internal"
	"femto-torrent/internal/torrent"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	fynedialog "fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	sqdialog "github.com/sqweek/dialog"
)

func main() {
	app := internal.NewApp()

	magnetInput := widget.NewEntry()
	magnetInput.SetPlaceHolder("Paste magnet link...")

	status := widget.NewLabel("Idling...")
	torrentList := container.NewVBox(
		widget.NewLabel("Ubuntu ISO"),
		widget.NewProgressBar(),
	)

	addButton := widget.NewButton("Add", func() {
		// TODO: Implement direct magnet string adding logic here
	})

	importButton := widget.NewButton("Import from file", func() {
		go handleFileImport(app.MainWindow, torrentList)
	})

	buttonRow := container.NewHBox(addButton, importButton)
	top := container.NewBorder(nil, nil, nil, buttonRow, magnetInput)
	content := container.NewBorder(top, status, nil, nil, container.NewVScroll(torrentList))

	app.MainWindow.SetContent(content)
	app.MainWindow.SetFixedSize(true)
	app.MainWindow.Resize(fyne.NewSize(600, 400))
	app.MainWindow.ShowAndRun()
}

func handleFileImport(w fyne.Window, torrentList *fyne.Container) {
	filename, err := sqdialog.File().
		Filter("Torrent files", "torrent", "magnet").
		Load()

	if err != nil {
		return
	}

	switch filepath.Ext(filename) {
	case ".torrent":
		t, err := torrent.DecodeTorrent(filename)
		if err != nil {
			fynedialog.ShowError(err, w)
			return
		}

		internal.ShowFileSelect(t, w, func(selected []internal.SelectedFile) {
			torrentList.Add(widget.NewLabel(t.Info.Name))
			torrentList.Add(widget.NewProgressBar())
			torrentList.Refresh()
		})

	case ".magnet":
		data, err := os.ReadFile(filename)
		if err != nil {
			fynedialog.ShowError(err, w)
			return
		}

		m, err := torrent.DecodeMagnet(strings.TrimSpace(string(data)))
		if err != nil {
			fynedialog.ShowError(err, w)
			return
		}

		fynedialog.ShowInformation("Magnet loaded!", "Name: "+m.Name, w)

	default:
		fynedialog.ShowError(errors.New("Unsupported file extension: "+filepath.Ext(filename)), w)
	}
}

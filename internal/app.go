package internal

import (
	"femto-torrent/internal/torrent"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	FyneApp     fyne.App
	MainWindow  fyne.Window
	StatusLabel *widget.Label
	ListElement *fyne.Container
	TorrentList []*torrent.Torrent
}

func NewApp() *App {
	fApp := app.New()
	w := fApp.NewWindow("Femto Torrent")

	status := widget.NewLabel("Idling...")
	list := container.NewVBox()

	return &App{
		FyneApp:     fApp,
		MainWindow:  w,
		StatusLabel: status,
		ListElement: list,
		TorrentList: []*torrent.Torrent{},
	}
}

func (a *App) AddTorrent(t *torrent.Torrent) {
	a.TorrentList = append(a.TorrentList, t)

	a.ListElement.Add(widget.NewLabel(t.Info.Name))
	a.ListElement.Add(widget.NewProgressBar())
	a.ListElement.Refresh()
}

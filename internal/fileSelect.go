package internal

import (
	"femto-torrent/internal/torrent"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type SelectedFile struct {
	Name string
	Size int64
}

type FileItem struct {
	Name     string
	Size     int64
	Enabled  bool
	Checkbox *widget.Check
}

func ShowFileSelect(t *torrent.Torrent, parent fyne.Window, onStart func(selected []SelectedFile)) {
	a := fyne.CurrentApp()
	w2 := a.NewWindow("Femto Torrent : " + t.Info.Name)

	var items []*FileItem

	if len(t.Info.Files) > 0 {
		for _, f := range t.Info.Files {
			name := strings.Join(f.Path, "/")

			item := &FileItem{
				Name:    name,
				Size:    f.Length,
				Enabled: true,
			}

			item.Checkbox = widget.NewCheck(name, func(checked bool) {
				item.Enabled = checked
			})

			item.Checkbox.SetChecked(true)

			items = append(items, item)
		}
	} else {
		item := &FileItem{
			Name:    t.Info.Name,
			Size:    t.Info.Length,
			Enabled: true,
		}

		item.Checkbox = widget.NewCheck(t.Info.Name, func(checked bool) {
			item.Enabled = checked
		})

		item.Checkbox.SetChecked(true)

		items = append(items, item)
	}

	content := container.NewVBox()

	for _, item := range items {
		content.Add(item.Checkbox)
	}

	w2StartButton := widget.NewButton("Start", func() {
		var selected []SelectedFile

		for _, item := range items {
			if item.Enabled {
				selected = append(selected, SelectedFile{
					Name: item.Name,
					Size: item.Size,
				})
			}
		}

		onStart(selected)

		w2.Close()
	})

	w2ButtonRow := container.NewHBox(
		w2StartButton,
	)

	w2Contents := container.NewBorder(
		nil,
		w2ButtonRow,
		nil,
		nil,
		content,
	)

	w2.SetContent(w2Contents)

	w2.Resize(fyne.NewSize(400, 300))
	w2.Show()
}

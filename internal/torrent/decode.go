package torrent

import (
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/jackpal/bencode-go"
)

type Torrent struct {
	Announce string `bencode:"announce"`
	Info     Info   `bencode:"info"`
}

type Magnet struct {
	InfoHash string
	Name     string
	Trackers []string
}

type Info struct {
	Name        string `bencode:"name"`
	PieceLength int    `bencode:"piece length"`
	Pieces      string `bencode:"pieces"`

	Length int64 `bencode:"length"`

	Files []struct {
		Length int64    `bencode:"length"`
		Path   []string `bencode:"path"`
	} `bencode:"files"`
}

func DecodeMagnet(link string) (*Magnet, error) {
	u, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	if u.Scheme != "magnet" {
		return nil, fmt.Errorf("not a magnet link")
	}

	q := u.Query()

	infoHash := strings.TrimPrefix(q.Get("xt"), "urn:btih:")

	return &Magnet{
		InfoHash: infoHash,
		Name:     q.Get("dn"),
		Trackers: q["tr"],
	}, nil
}

func DecodeTorrent(path string) (*Torrent, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var torrent Torrent
	err = bencode.Unmarshal(f, &torrent)
	if err != nil {
		return nil, err
	}

	return &torrent, nil
}

func (t *Torrent) FileList() []string {
	var files []string

	if len(t.Info.Files) > 0 {
		for _, f := range t.Info.Files {
			path := strings.Join(f.Path, "/")
			files = append(files, fmt.Sprintf("%s (%d bytes)", path, f.Length))
		}
		return files
	}

	files = append(files, fmt.Sprintf("%s (%d bytes)", t.Info.Name, t.Info.Length))
	return files
}

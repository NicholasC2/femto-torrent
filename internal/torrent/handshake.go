package torrent

type Handshake struct {
	Pstr     string
	InfoHash [20]byte
	PeerID   [20]byte
}

func (h *Handshake) Serialize() []byte {
	buf := make([]byte, 49+len(h.Pstr))
	buf[0] = byte(len(h.Pstr))

	offset := 1
	copy(buf[offset:], h.Pstr)
	offset += len(h.Pstr)

	for i := 0; i < 8; i++ {
		buf[offset+i] = 0
	}
	offset += 8

	copy(buf[offset:], h.InfoHash[:])
	offset += 20

	copy(buf[offset:], h.PeerID[:])
	return buf
}

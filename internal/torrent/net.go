package torrent

import (
	"io"
	"net"
)

func DoHandshake(conn net.Conn, infoHash, peerID [20]byte) ([20]byte, error) {
	hs := Handshake{
		Pstr:     ProtocolName,
		InfoHash: infoHash,
		PeerID:   peerID,
	}

	_, err := conn.Write(hs.Serialize())
	if err != nil {
		return [20]byte{}, err
	}

	resp := make([]byte, HandshakeLen)
	_, err = io.ReadFull(conn, resp)
	if err != nil {
		return [20]byte{}, err
	}

	var peerInfoHash [20]byte
	copy(peerInfoHash[:], resp[28:48])

	return peerInfoHash, nil
}

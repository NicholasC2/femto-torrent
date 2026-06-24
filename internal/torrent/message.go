package torrent

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

type Message struct {
	ID      byte
	Payload []byte
}

func ReadMessage(conn net.Conn) (*Message, error) {
	var length uint32
	err := binary.Read(conn, binary.BigEndian, &length)
	if err != nil {
		return nil, err
	}

	// keep-alive
	if length == 0 {
		return nil, nil
	}

	msg := make([]byte, length)
	_, err = io.ReadFull(conn, msg)
	if err != nil {
		return nil, err
	}

	return &Message{
		ID:      msg[0],
		Payload: msg[1:],
	}, nil
}

func SendMessage(conn net.Conn, id byte, payload []byte) error {
	var buf bytes.Buffer

	length := uint32(len(payload) + 1)
	if err := binary.Write(&buf, binary.BigEndian, length); err != nil {
		return err
	}

	buf.WriteByte(id)
	buf.Write(payload)

	_, err := conn.Write(buf.Bytes())
	return err
}

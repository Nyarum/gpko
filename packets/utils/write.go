package utils

import (
	"encoding/binary"
	"io"
)

func WritePKO(w io.Writer, order binary.ByteOrder, data any) error {
	switch v := data.(type) {
	case string:
		err := binary.Write(w, binary.BigEndian, uint16(len(v)))
		if err != nil {
			return err
		}

		return binary.Write(w, order, []byte(v))
	default:
		return binary.Write(w, order, data)
	}
}

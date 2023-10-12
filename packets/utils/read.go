package utils

import (
	"encoding/binary"
	"io"
)

func ReadPKO(r io.Reader, order binary.ByteOrder, data any) error {
	switch v := data.(type) {
	case *string:
		var ln uint16
		err := binary.Read(r, binary.BigEndian, &ln)
		if err != nil {
			return err
		}

		buf := make([]byte, ln)
		_, err = r.Read(buf)
		if err != nil {
			return err
		}

		*v = string(buf)

		return nil
	default:
		return binary.Read(r, order, v)
	}
}

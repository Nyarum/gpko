package packets

import (
	"bytes"
	"encoding/binary"

	"github.com/samber/do"
)

type Outcome interface {
	Write() ([]byte, error)
	Opcode() Opcode
}

type Builder struct {
}

func NewBuilder(i *do.Injector) (Builder, error) {
	return Builder{}, nil
}

func (b Builder) Build(out Outcome) ([]byte, error) {
	data, err := out.Write()
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer([]byte{})
	err = binary.Write(buf, binary.BigEndian, uint16(len(data)+6))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, uint32(128))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

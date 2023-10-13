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

type Income interface {
	Read(in []byte) error
	Handle()
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
	err = binary.Write(buf, binary.BigEndian, uint16(len(data)+8))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, uint32(128))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, out.Opcode())
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (b Builder) Unbuild(in []byte) (Income, error) {
	buf := bytes.NewBuffer(in)

	var (
		ln     uint16
		id     uint32
		opcode uint16
	)

	err := binary.Read(buf, binary.BigEndian, &ln)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.LittleEndian, &id)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.BigEndian, &opcode)
	if err != nil {
		return nil, err
	}

	income := SwitchInPackets[Opcode(opcode)]()
	err = income.Read(buf.Next(int(ln - 8)))
	if err != nil {
		return nil, err
	}

	return income, nil
}

package out

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"gpko/packets"
	"gpko/packets/utils"
	"time"
)

type Date struct {
}

func NewDate() Date {
	return Date{}
}

func (d Date) Opcode() packets.Opcode {
	return packets.OP_SERVER_CHAPSTR
}

func (d Date) Write() ([]byte, error) {
	buf := bytes.NewBuffer([]byte{})

	timeNow := time.Now()
	output := fmt.Sprintf(
		"[%02d-%02d %02d:%02d:%02d:%03d]",
		timeNow.Month(), timeNow.Day(),
		timeNow.Hour(), timeNow.Minute(), timeNow.Second(),
		timeNow.Nanosecond()/1000000,
	)

	err := utils.WritePKO(buf, binary.LittleEndian, output)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

package utils

import (
	"gpko/packets"
)

type Income interface {
	Read()
	Handle()
}

type Outcome interface {
	Write() ([]byte, error)
}

var (
	SwitchInPackets = map[packets.Opcode]Income{}
)

func init() {
	//outDate := out.Date{}
	//SwitchInPackets[outDate.Opcode()] = outDate
}

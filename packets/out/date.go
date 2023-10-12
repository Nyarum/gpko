package out

import (
	"fmt"
	"gpko/packets"
	"time"
)

type Date struct {
}

func NewDate() Date {
	return Date{}
}

func (d Date) Opcode() packets.Opcode {
	return packets.OP_SERVER_LOGIN
}

func (d Date) Write() ([]byte, error) {
	timeNow := time.Now()
	output := fmt.Sprintf(
		"[%02d-%02d %02d:%02d:%02d:%03d]",
		timeNow.Month(), timeNow.Day(),
		timeNow.Hour(), timeNow.Minute(), timeNow.Second(),
		timeNow.Nanosecond()/1000000,
	)

	return []byte(output), nil
}

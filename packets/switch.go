package packets

import "gpko/packets/in"

var (
	SwitchInPackets = map[Opcode]func() Income{
		OP_CLIENT_LOGIN: func() Income {
			return &in.Auth{}
		},
	}
)

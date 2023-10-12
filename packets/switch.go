package packets

import "gpko/packets/in"

var (
	SwitchInPackets = map[Opcode]interface{}{
		OP_CLIENT_LOGIN: in.Auth{},
	}
)

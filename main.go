package main

import (
	"gpko/packets"
	"gpko/server"

	"github.com/samber/do"
	"golang.org/x/exp/slog"
)

func main() {
	injector := do.New()

	do.Provide[server.TCP](injector, server.NewTCP)
	do.Provide[server.Client](injector, server.NewClient)
	do.Provide[packets.Builder](injector, packets.NewBuilder)

	tcp := do.MustInvoke[server.TCP](injector)
	err := tcp.Listen(":1973")
	if err != nil {
		slog.Error(err.Error())
		return
	}
}

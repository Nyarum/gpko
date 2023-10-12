package server

import (
	"fmt"
	"gpko/packets"
	"gpko/packets/out"
	"net"

	"github.com/samber/do"
	"golang.org/x/exp/slog"
)

type Client struct {
	builder packets.Builder
}

func NewClient(i *do.Injector) (Client, error) {
	builder := do.MustInvoke[packets.Builder](i)

	return Client{
		builder: builder,
	}, nil
}

func (client Client) Listen(c net.Conn) error {
	buf := make([]byte, 4092)
	for {
		res, err := client.builder.Build(out.NewDate())
		if err != nil {
			slog.Error("failed to build first packet", "error", err)
			return err
		}

		fmt.Printf("% x\n", res)

		ln, err := c.Read(buf)
		if err != nil {
			slog.Error("failed to read from connection", "error", err)
			return err
		}

		slog.Info("has read from connection", "len", ln)
		slog.Info("buf", "bytes", buf)
	}
}

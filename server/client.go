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
	return Client{
		builder: do.MustInvoke[packets.Builder](i),
	}, nil
}

func (client Client) Listen(c net.Conn) error {
	res, err := client.builder.Build(
		out.NewDate(),
	)
	if err != nil {
		return err
	}

	fmt.Printf("% x\n", res)

	_, err = c.Write(res)
	if err != nil {
		return err
	}

	buf := make([]byte, 4092)
	for {
		ln, err := c.Read(buf)
		if err != nil {
			return err
		}

		if ln == 2 {
			c.Write([]byte{0x00, 0x02})
			continue
		}

		slog.Info("has read from connection", "len", ln)
		slog.Info("buf", "bytes", buf[:ln])

		income, err := client.builder.Unbuild(buf[:ln])
		if err != nil {
			return err
		}

		income.Handle()
	}
}

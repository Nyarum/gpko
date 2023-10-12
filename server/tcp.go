package server

import (
	"net"

	"github.com/samber/do"
	"golang.org/x/exp/slog"
)

type TCP struct {
	i *do.Injector
}

func NewTCP(i *do.Injector) (TCP, error) {
	return TCP{
		i: i,
	}, nil
}

func (tcp *TCP) Listen(host string) error {
	l, err := net.Listen("tcp4", host)
	if err != nil {
		return err
	}
	defer l.Close()

	slog.Info("has running tcp server", "host", host)

	for {
		c, err := l.Accept()
		if err != nil {
			slog.Error("failed to accept connection", "error", err)
			return err
		}

		h := do.MustInvoke[Client](tcp.i)

		go func() {
			if err := h.Listen(c); err != nil {
				slog.Error("handle client returned error", "error", err)
			}
		}()
	}
}

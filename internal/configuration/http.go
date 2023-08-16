package configuration

import "time"

type HTTP struct {
	address         string
	shutdownTimeout time.Duration
}

func (h HTTP) Address() string {
	return h.address
}

func (h HTTP) ShutdownTimeout() time.Duration {
	return h.shutdownTimeout
}

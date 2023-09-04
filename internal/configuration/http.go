package configuration

import (
	"fmt"
	"time"
)

// HTTP is the configuration used to set the http server instance.
type HTTP struct {
	host            string
	port            int
	shutdownTimeout time.Duration
}

// Address returns the TCP address for the server to listen on, in the form "host:port".
func (h HTTP) Address() string {
	return fmt.Sprintf("%v:%d", h.host, h.port)
}

// ShutdownTimeout returns the timeout duration used on shutting the server down.
func (h HTTP) ShutdownTimeout() time.Duration {
	return h.shutdownTimeout
}

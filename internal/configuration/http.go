package configuration

import "time"

// HTTP is the configuration used to set the http server instance.
type HTTP struct {
	address         string
	shutdownTimeout time.Duration
}

// Address returns the TCP address for the server to listen on, in the form "host:port".
func (h HTTP) Address() string {
	return h.address
}

// ShutdownTimeout returns the timeout duration used on shutting the server down.
func (h HTTP) ShutdownTimeout() time.Duration {
	return h.shutdownTimeout
}

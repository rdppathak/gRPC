//
// Author: rdppathak@gmail.com
//
// Server configuration related functionality

package server

// ServerConfig to hold configuration details
type serverConfig struct {
	port    int
	address string
}

func NewServerConfig(address string, port int) *serverConfig {
	return &serverConfig{
		port:    port,
		address: address,
	}
}

// GetPort returns the port number for given configuration
func (sc *serverConfig) GetPort() int {
	return sc.port
}

// GetAddress returns the address on which server will bind
func (sc *serverConfig) GetAddress() string {
	return sc.address
}

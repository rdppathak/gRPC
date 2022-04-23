//
// Author: rdppathak@gmail.com
//
// Server package for handlers

package server

import (
	"fmt"
	"net/http"

	"github.com/golang/glog"
)

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

// Serve
func Serve(serverConfig *serverConfig) {
	glog.Infof("Initialising server with config: %v", serverConfig)

	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/rpc", rpcHandler)

	addressAndPort := fmt.Sprintf("%s:%d",
		serverConfig.address, serverConfig.port)
	http.ListenAndServe(addressAndPort, nil)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Default handler invoked...")
	fmt.Fprintf(w, "Defaulthandler\n")
}

func rpcHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("RPC handler function...")
	fmt.Fprintf(w, "RPCHandler\n")
}

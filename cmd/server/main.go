//
// Author: rdppathak@gmail.com
//
// gRPC server entry point

package main

import (
	"flag"

	"github.com/golang/glog"
	"github.com/rdppathak/gRPC/pkg/server"
)

func main() {
	// Parse flags
	flag.Parse()

	// TODO: Set logging level and configuration
	glog.Infof("Starting gRPC server...")

	// TODO: read from config file
	serverConfig := server.NewServerConfig("127.0.0.1", 8080)

	server.Serve(serverConfig)
}

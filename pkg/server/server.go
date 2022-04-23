//
// Author: rdppathak@gmail.com
//
// Server package for handlers

package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/glog"
)

type rpcGenericRequest struct {
	MethodName string      `json:"method_name"`
	Args       interface{} `json:"args"`
}

type server struct {
	serverConfig   *serverConfig
	registeredRPCs map[string]func(args interface{})
}

func NewServer(serverConfig *serverConfig) *server {
	return &server{
		serverConfig:   serverConfig,
		registeredRPCs: make(map[string]func(args interface{})),
	}
}

// Serve
func (s *server) Serve() {
	glog.Infof("Initialising server with config: %v", s.serverConfig)

	http.HandleFunc("/", s.defaultHandler)
	http.HandleFunc("/rpc", s.rpcHandler)

	addressAndPort := fmt.Sprintf("%s:%d",
		s.serverConfig.GetAddress(), s.serverConfig.GetPort())
	http.ListenAndServe(addressAndPort, nil)
}

func (s *server) RegisterRPC(methodName string, handler func(interface{})) error {
	if _, ok := s.registeredRPCs[methodName]; ok {
		return fmt.Errorf("Method %s already registered..", methodName)
	}

	s.registeredRPCs[methodName] = handler
	return nil
}

func (s *server) defaultHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("Default handler invoked...")
	fmt.Fprintf(w, "Defaulthandler\n")
}

func (s *server) rpcHandler(w http.ResponseWriter, r *http.Request) {
	glog.Infof("RPC handler function...")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Errorf("Failed to read request body: %s", err.Error())
		w.WriteHeader(500)
		return
	}

	var rpcRequest rpcGenericRequest
	if err = json.Unmarshal(body, &rpcRequest); err != nil {
		glog.Errorf("Body parse error, %s", err.Error())
		w.WriteHeader(http.StatusBadRequest) // Return 400 Bad Request.
		return
	}

	if glog.V(2) {
		glog.Infof("Received the RPC request: %s", body)
	}

	rpcFunc, ok := s.registeredRPCs[rpcRequest.MethodName]
	if !ok {
		glog.Errorf("RPC [%s] not found", rpcRequest.MethodName)
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, fmt.Sprintf("RPC [%s] not found\n",
			rpcRequest.MethodName))
		return
	}

	rpcFunc(rpcRequest.Args)
	fmt.Fprintf(w, "RPCHandler\n")
}

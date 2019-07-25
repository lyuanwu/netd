package ingress

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"

	"github.com/songtianyi/rrframework/logs"
)

// Jrpc struct
type Jrpc struct {
	addr string
}

// NewJrpc return created jrpc instance
func NewJrpc(addr string) (*Jrpc, error) {
	return &Jrpc{addr: addr}, nil
}

// Register do jrpc logic method registration
func (s *Jrpc) Register(method interface{}) error {
	return rpc.Register(method)
}

// Serve start listen tcp port and accept jrpc calls
func (s *Jrpc) Serve() error {
	listener, e := net.Listen("tcp", s.addr)
	if e != nil {
		return e
	}
	for {
		if conn, err := listener.Accept(); err != nil {
			logs.Error("accept error: " + err.Error())
		} else {
			logs.Info("new jrpc connection established")
			go jsonrpc.ServeConn(conn)
		}
	}
}

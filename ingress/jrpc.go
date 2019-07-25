// NetD makes network device operations easy.
// Copyright (C) 2019  sky-cloud.net
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

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

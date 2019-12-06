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
	"time"

	"github.com/sky-cloud-tec/netd/common"
	"github.com/sky-cloud-tec/netd/protocol"
	"github.com/songtianyi/rrframework/logs"
	"github.com/songtianyi/rrframework/utils"
)

// UtilsHandler serve many useful services like port checking
type UtilsHandler struct {
}

// CheckPort check specifed port is open or not
func (s *UtilsHandler) CheckPort(req *protocol.PortCheckRequest, res *protocol.PortCheckResponse) error {
	logs.Info("Receiving req", req)

	// build timeout
	if req.Timeout == 0 {
		req.Timeout = common.DefaultTimeout
	} else {
		req.Timeout = req.Timeout * time.Second
	}

	// build log prefix
	if req.LogPrefix == "" {
		req.LogPrefix = "[ " + req.Proto + "://" + req.IP + ":" + req.Port + " ]"
	}

	// build session
	if req.Session == "" {
		req.Session = rrutils.NewV4().String()
	}
	req.LogPrefix = req.LogPrefix + " [ " + req.Session + " ] "

	ch := make(chan error, 1)

	go func() {
		logs.Info(req.LogPrefix, "==========START==========")
		logs.Info(req.LogPrefix, "==========END==========")
	}()

	// timeout
	select {
	case res := <-ch:
		return res
	case <-time.After(req.Timeout):
		*res = makeUtilsErrRes(common.ErrNoOpFound, "handle req timeout")
	}

	return nil
}

func makeUtilsErrRes(code int, msg string) protocol.PortCheckResponse {
	return protocol.PortCheckResponse{Retcode: code, Message: msg}
}

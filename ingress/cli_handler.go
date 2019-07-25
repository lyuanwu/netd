package ingress

import (
	"strings"
	"time"

	"github.com/sky-cloud-tec/netd/cli"
	"github.com/sky-cloud-tec/netd/cli/conn"
	_ "github.com/sky-cloud-tec/netd/cli/juniper/srx" // load juniper srx
	"github.com/sky-cloud-tec/netd/common"
	"github.com/sky-cloud-tec/netd/protocol"
	"github.com/songtianyi/rrframework/logs"
)

// CliHandler run cli commands and return result to caller
type CliHandler struct {
}

// Handle cli request
func (s *CliHandler) Handle(req *protocol.CliRequest, res *protocol.CliResponse) error {
	logs.Info("Receiving req", req)
	// build timeout
	if req.Timeout == 0 {
		req.Timeout = common.DefaultTimeout
	} else {
		req.Timeout = req.Timeout * time.Second
	}

	// build log prefix
	if req.LogPrefix == "" {
		req.LogPrefix = "[ " + req.Device + " ]"
	}

	ch := make(chan error, 1)

	go func() {
		ch <- doHandle(req, res)
	}()

	// timeout
	select {
	case res := <-ch:
		return res
	case <-time.After(req.Timeout):
		*res = makeCliErrRes(common.ErrNoOpFound, "handle req timeout")
	}
	return nil
}

func doHandle(req *protocol.CliRequest, res *protocol.CliResponse) error {
	// build device operator type
	t := strings.Join([]string{req.Vendor, req.Type, req.Version}, ".")
	// get operator by type
	op := cli.OperatorManagerInstance.Get(t)
	if op == nil {
		logs.Error(req.LogPrefix, "no operator match", t)
		*res = makeCliErrRes(common.ErrNoOpFound, "no operator match "+t)
		return nil
	}
	// acquire cli connection, it could be blocked here for concurrency
	c, err := conn.Acquire(req, op)
	defer conn.Release(req)
	if err != nil {
		logs.Error(req.LogPrefix, "new operator fail,", err)
		*res = makeCliErrRes(common.ErrNewOp, "new operator fail, "+err.Error())
		return nil
	}
	// execute cli commands
	out, err := c.Exec()
	if err != nil {
		logs.Error(req.LogPrefix, "exec error,", err)
		*res = makeCliErrRes(common.ErrOpExec, "exec cli cmds fail, "+err.Error())
		return nil
	}
	// make reponse
	*res = protocol.CliResponse{
		Retcode: common.OK,
		Message: "OK",
		CmdsStd: out,
	}
	return nil
}

func makeCliErrRes(code int, msg string) protocol.CliResponse {
	return protocol.CliResponse{Retcode: code, Message: msg, CmdsStd: nil}
}

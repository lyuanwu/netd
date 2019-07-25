package ingress

import (
	"net"
	"net/rpc/jsonrpc"
	"testing"

	"github.com/sky-cloud-tec/netd/common"
	"github.com/sky-cloud-tec/netd/protocol"

	. "github.com/smartystreets/goconvey/convey"
)

func TestJuniperSrx_Set(t *testing.T) {

	Convey("set juniper srx cli commands", t, func() {
		client, err := net.Dial("tcp", "localhost:8088")
		So(
			err,
			ShouldBeNil,
		)
		// Synchronous call
		args := &protocol.CliRequest{
			Device:  "juniper-test",
			Vendor:  "juniper",
			Type:    "srx",
			Version: "6.0",
			Address: "192.168.1.252:22",
			Auth: protocol.Auth{
				Username: "admin",
				Password: "r00tme",
			},
			Commands: []string{"set security address-book global address WS-100.2.2.46_32 wildcard-address 100.2.2.46/32"},
			Protocol: "ssh",
			Mode:     "configure_private",
			Timeout:  30,
		}
		var reply protocol.CliResponse
		c := jsonrpc.NewClient(client)
		err = c.Call("CliHandler.Handle", args, &reply)
		So(
			err,
			ShouldBeNil,
		)
		So(
			reply.Retcode == common.OK,
			ShouldBeTrue,
		)
	})
}

func TestJuniperSrx_Show(t *testing.T) {

	Convey("show juniper srx configuration", t, func() {
		client, err := net.Dial("tcp", "localhost:8088")
		So(
			err,
			ShouldBeNil,
		)
		// Synchronous call
		args := &protocol.CliRequest{
			Device:  "juniper-test-show",
			Vendor:  "juniper",
			Type:    "srx",
			Version: "6.0",
			Address: "192.168.1.252:22",
			Auth: protocol.Auth{
				Username: "admin",
				Password: "r00tme",
			},
			Commands: []string{"show configuration | display set | no-more"},
			Protocol: "ssh",
			Mode:     "login",
			Timeout:  30,
		}
		var reply protocol.CliResponse
		c := jsonrpc.NewClient(client)
		err = c.Call("CliHandler.Handle", args, &reply)
		So(
			err,
			ShouldBeNil,
		)
		So(
			reply.Retcode == common.OK,
			ShouldBeTrue,
		)
		So(
			reply.CmdsStd,
			ShouldNotBeNil,
		)
		So(
			len(reply.CmdsStd) == 1,
			ShouldBeTrue,
		)
	})
}

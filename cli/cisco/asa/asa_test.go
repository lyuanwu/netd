package asa

import (
	"testing"

	"github.com/sky-cloud-tec/netd/cli"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAsaOp(t *testing.T) {

	Convey("asa op", t, func() {
		op := createOp9xPlus()
		So(
			cli.Match(op.GetPrompts("login"), "asaNAT> "),
			ShouldBeTrue,
		)
	})
}

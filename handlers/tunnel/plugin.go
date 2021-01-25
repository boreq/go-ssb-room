package tunnel

import (
	"net"

	kitlog "github.com/go-kit/kit/log"
	"go.cryptoscope.co/muxrpc/v2"
	"go.cryptoscope.co/muxrpc/v2/typemux"

	refs "go.mindeco.de/ssb-refs"
	"go.mindeco.de/ssb-rooms/internal/maybemuxrpc"
)

const name = "tunnel"

var method muxrpc.Method = muxrpc.Method{name}

type plugin struct {
	h   muxrpc.Handler
	log kitlog.Logger
}

func (plugin) Name() string              { return name }
func (plugin) Method() muxrpc.Method     { return method }
func (p plugin) Handler() muxrpc.Handler { return p.h }
func (plugin) Authorize(net.Conn) bool   { return true }

/* manifest:
{
	"announce": "sync",
	"leave": "sync",
	"connect": "duplex",
	"endpoints": "source",
	"isRoom": "async",
	"ping": "sync",
}
*/

func New(log kitlog.Logger, self refs.FeedRef) maybemuxrpc.Plugin {
	mux := typemux.New(log)

	var rs roomState
	rs.logger = log

	mux.RegisterAsync(append(method, "isRoom"), typemux.AsyncFunc(rs.isRoom))
	mux.RegisterAsync(append(method, "ping"), typemux.AsyncFunc(rs.ping))

	mux.RegisterAsync(append(method, "announce"), typemux.AsyncFunc(rs.announce))
	mux.RegisterAsync(append(method, "leave"), typemux.AsyncFunc(rs.leave))

	mux.RegisterSource(append(method, "endpoints"), typemux.SourceFunc(rs.endpoints))

	// TODO: patch typemux
	// mux.RegisterDuplex(append(method, "connect"), typemux.DuplexFunc(rs.connect))

	return plugin{
		h: &mux,
	}
}
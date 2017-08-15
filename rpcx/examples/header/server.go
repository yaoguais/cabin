package main

import (
	"context"
	"net"

	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/core"
	"github.com/smallnest/rpcx/log"
)

type Args struct {
	A   int
	B   int
	ctx map[string]interface{}
}

type Reply struct {
	C int
}

func (a *Args) Value(key string) interface{} {
	if a.ctx != nil {
		return a.ctx[key]
	}
	return nil
}

func (a *Args) SetValue(key string, value interface{}) {
	if a.ctx == nil {
		a.ctx = make(map[string]interface{})
	}
	a.ctx[key] = value
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	m, ok := core.FromMapContext(ctx)
	if ok {
		conn := m[core.ConnKey].(net.Conn)
		log.Infof("Client IP: %s", conn.RemoteAddr().String())
		header := m[core.HeaderKey].(core.Header)
		log.Infof("header: %s", header.String())
	}
	reply.C = args.A * args.B

	return nil
}

func main() {
	server := rpcx.NewServer()
	server.RegisterName("Arith", new(Arith))
	server.Serve("tcp", "127.0.0.1:8972")
}

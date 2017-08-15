package main

import (
	"context"

	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	server := rpcx.NewServer()
	server.RegisterName("Arith", new(Arith))

	p := plugin.NewCompressionPlugin(rpcx.CompressSnappy)
	server.PluginContainer.Add(p)

	server.Serve("tcp", "127.0.0.1:8972")
}

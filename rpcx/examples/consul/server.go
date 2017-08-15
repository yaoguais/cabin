package main

import (
	"context"
	"time"

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
	p := &plugin.ConsulRegisterPlugin{
		ServiceAddress: "tcp@127.0.0.1:8972",
		ConsulAddress:  "localhost:8500",
		UpdateInterval: time.Second,
	}
	p.Start()
	server.PluginContainer.Add(p)

	server.RegisterName("Arith", new(Arith), "weight=5&group=beijing")

	server.Serve("tcp", "127.0.0.1:8972")
}

package main

import (
	"context"
	"time"

	"github.com/smallnest/rpcx"
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
	time.Sleep(time.Minute)
	return nil
}

func main() {
	server := rpcx.NewServer()
	server.Timeout = 120 * time.Second
	server.ReadTimeout = 120 * time.Second
	server.WriteTimeout = 120 * time.Second

	server.RegisterName("Arith", new(Arith))
	server.Serve("tcp", "127.0.0.1:8972")
}

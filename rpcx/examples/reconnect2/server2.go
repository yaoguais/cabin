package main

import (
	"context"

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
	return nil
}

type Arith2 int

func (t *Arith2) Mul(args *Args, reply *Reply) error {
	reply.C = args.A * args.B * 10
	return nil
}

func main() {
	server2 := rpcx.NewServer()
	server2.RegisterName("Arith", new(Arith2))
	server2.Serve("tcp", "127.0.0.1:8973")
}

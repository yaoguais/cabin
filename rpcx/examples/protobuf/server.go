package main

import (
	"context"
	"time"

	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/codec"
	"github.com/smallnest/rpcx/plugin"
)

type Arith int32

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

func main() {
	server := rpcx.NewServer()
	server.ServerCodecFunc = codec.NewProtobufServerCodec
	server.RegisterName("Arith", new(Arith))

	p := plugin.NewRateLimitingPlugin(time.Second, 1000)
	server.PluginContainer.Add(p)

	server.Serve("tcp", "127.0.0.1:8972")
}

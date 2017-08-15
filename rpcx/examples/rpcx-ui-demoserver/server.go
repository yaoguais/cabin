package main

import (
	"context"
	"flag"
	"time"

	"github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/plugin"
)

type Args struct {
	A int `msg:"a"`
	B int `msg:"b"`
}

type Reply struct {
	C int `msg:"c"`
}

type Arith int

func (t *Arith) Mul(ctx context.Context, args *Args, reply *Reply) error {
	reply.C = args.A * args.B
	return nil
}

var addr = flag.String("s", "127.0.0.1:8972", "service address")
var zk = flag.String("zk", "127.0.0.1:2181", "zookeeper URL")
var n = flag.String("n", "Arith", "Service Name")
var g = flag.String("g", "", "Group Name")

// start multiple servers and check rpcx-ui
func main() {
	flag.Parse()

	server := rpcx.NewServer()
	rplugin := &plugin.ZooKeeperRegisterPlugin{
		ServiceAddress:   "tcp@" + *addr,
		ZooKeeperServers: []string{*zk},
		BasePath:         "/rpcx",
		Metrics:          metrics.NewRegistry(),
		Services:         make([]string, 0),
		UpdateInterval:   10 * time.Second,
	}
	server.PluginContainer.Add(rplugin)

	rplugin.Start()
	server.RegisterName(*n, new(Arith), "weight=5&state=active&group="+*g)
	server.Serve("tcp", *addr)
}

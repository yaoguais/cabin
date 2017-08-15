package main

import (
	"context"
	"time"

	"github.com/apex/log"
	"github.com/smallnest/rpcx"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	s := &rpcx.DirectClientSelector{Network: "tcp", Address: "127.0.0.1:8972", DialTimeout: 10 * time.Second}
	client := rpcx.NewClient(s)
	client.Timeout = 10 * time.Second
	client.ReadTimeout = 10 * time.Second
	client.WriteTimeout = 10 * time.Second

	args := &Args{7, 8}
	var reply Reply
	err := client.Call(context.Background(), "Arith.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	} else {
		log.Infof("Arith: %d*%d=%d", args.A, args.B, reply.C)
	}

	client.Close()
}

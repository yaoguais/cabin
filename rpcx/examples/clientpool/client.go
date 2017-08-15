package main

import (
	"context"
	"sync"
	"time"

	"github.com/smallnest/pool"
	"github.com/smallnest/rpcx"
	"github.com/smallnest/rpcx/clientselector"
	"github.com/smallnest/rpcx/log"
)

type Args struct {
	A int
	B int
}

type Reply struct {
	C int
}

func main() {
	server1 := &clientselector.ServerPeer{Network: "tcp", Address: "127.0.0.1:8972"}
	server2 := &clientselector.ServerPeer{Network: "tcp", Address: "127.0.0.1:8973"}

	servers := []*clientselector.ServerPeer{server1, server2}

	s := clientselector.NewMultiClientSelector(servers, rpcx.RandomSelect, 10*time.Second)

	clientPool := &pool.Pool{
		New: func() interface{} {
			return rpcx.NewClient(s)
		},
	}

	var sg sync.WaitGroup
	sg.Add(1000)
	for i := 0; i < 1000; i++ {
		go callServer(&sg, clientPool, s)
		time.Sleep(10 * time.Millisecond)
	}

	sg.Wait()

	clientPool.Range(func(v interface{}) bool {
		c := v.(*rpcx.Client)
		c.Close()
		return true
	})
	clientPool.Reset()
}

func callServer(sg *sync.WaitGroup, clientPool *pool.Pool, s rpcx.ClientSelector) {
	client := clientPool.Get().(*rpcx.Client)

	args := &Args{7, 8}
	var reply Reply
	err := client.Call(context.Background(), "Arith.Mul", args, &reply)
	if err != nil {
		log.Infof("error for Arith: %d*%d, %v", args.A, args.B, err)
	} else {
		log.Infof("Arith: %d*%d=%d, client: %p", args.A, args.B, reply.C, client)
	}

	clientPool.Put(client)
	sg.Done()
}

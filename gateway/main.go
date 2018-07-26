package main

import (
	"flag"
	"fmt"
	"strconv"
	"time"

	"github.com/fy026/weather/lib/client"
	"github.com/fy026/weather/lib/registry"
	"github.com/fy026/weather/lib/registry/etcd"
	"github.com/fy026/weather/proto"
	"golang.org/x/net/context"
)

var (
	serv = flag.String("service", "ts", "service name")
	reg  = flag.String("reg", "http://127.0.0.1:2379", "register etcd address")
)

func main() {
	flag.Parse()

	registry, err := etcd.NewEtcdRegisty(registry.WithAddress(*reg))
	if err != nil {
		return
	}

	c := client.NewClient(
		client.WithServiceName(*serv),
		client.WithTimeout(time.Second*10),
		client.WithRegistry(registry),
	)

	grpcConn, err := c.GetConn()
	if err != nil {
		return
	}

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		client := pb.NewGreeterClient(grpcConn)
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world " + strconv.Itoa(t.Second())})
		if err == nil {
			fmt.Printf("%v: Reply is %s\n", t, resp.Message)
		} else {
			fmt.Println(err)
		}
	}

}

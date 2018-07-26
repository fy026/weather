package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fy026/weather/examples/proto"
	"github.com/fy026/weather/lib/registry"
	"github.com/fy026/weather/lib/registry/etcd"
	"github.com/fy026/weather/lib/server"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

var (
	serv = flag.String("n", "ts", "service name")
	host = flag.String("h", "127.0.0.1", "listening port")
	port = flag.String("p", "50001", "listening port")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

type testServer struct{}

// SayHello implements helloworld.GreeterServer
func (s *testServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s\n", time.Now(), in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	flag.Parse()

	registry, err := etcd.NewEtcdRegisty(registry.WithAddress(*reg))
	if err != nil {
		return
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		fmt.Printf("receive signal '%v'", s)
		registry.Deregister()
		os.Exit(1)
	}()

	s := server.NewServer(
		server.WithServiceName(*serv),
		server.WithServiceId(uuid.NewUUID().String()),
		server.WithHost(*host),
		server.WithPort(*port),
		server.WithRegistry(registry),
	)
	if coreServ := s.GetGRPCServer(); coreServ != nil {
		pb.RegisterGreeterServer(coreServ, &testServer{})
		s.Start()
	}

}

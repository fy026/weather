package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fy026/weather/pkg/registry"
	"github.com/fy026/weather/pkg/registry/etcd"
	"github.com/fy026/weather/pkg/server"
	"github.com/fy026/weather/proto"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
)

var (
	serv = flag.String("n", "ts", "service name")
	host = flag.String("h", "127.0.0.1", "listening port")
	port = flag.String("p", "50001", "listening port")
	reg  = flag.String("reg", "http://localhost:2379", "register etcd address")
)

type testServer struct {
	ServerId string
}

// SayHello implements helloworld.GreeterServer
func (s *testServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("%v: Receive is %s service id :%s\n", time.Now(), in.Name, s.ServerId)
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %s %s", in.Name, s.ServerId)}, nil
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

	serviceId := uuid.NewUUID().String()
	s := server.NewServer(
		server.WithServiceName(*serv),
		server.WithServiceId(serviceId),
		server.WithHost(*host),
		server.WithPort(*port),
		server.WithRegistry(registry),
	)

	ts := testServer{ServerId: serviceId}

	if coreServ := s.GetGRPCServer(); coreServ != nil {
		pb.RegisterGreeterServer(coreServ, &ts)
		s.Start()
	}

}

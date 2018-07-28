package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fy026/weather/pkg/registry"
	"github.com/fy026/weather/pkg/registry/etcd"
	"github.com/fy026/weather/pkg/server"
	"github.com/fy026/weather/proto"
	"github.com/fy026/weather/service/pkg/setting"
	"github.com/pborman/uuid"
	"golang.org/x/net/context"
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
	serviceId := uuid.NewUUID().String()
	fmt.Printf("server service name:%s\n", setting.Server_Name)
	opts := []server.Options{
		server.WithServiceName(setting.Server_Name),
		server.WithServiceId(serviceId),
		server.WithHost(setting.HTTPHost),
		server.WithPort(setting.HTTPPort),
	}

	if setting.ENV != "k8s" {
		registry, err := etcd.NewEtcdRegisty(registry.WithAddress(setting.RegisterUrl))
		if err != nil {
			fmt.Println("registy error:", err.Error())
			return
		}
		opts = append(opts, server.WithRegistry(registry))
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
		go func() {
			s := <-ch
			fmt.Printf("receive signal '%v'", s)
			registry.Deregister()
			os.Exit(1)
		}()
	}

	s := server.NewServer(opts...)

	ts := testServer{ServerId: serviceId}

	if coreServ := s.GetGRPCServer(); coreServ != nil {
		pb.RegisterGreeterServer(coreServ, &ts)
		s.Start()
	}

}

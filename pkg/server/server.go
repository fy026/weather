package server

import (
	"context"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/naming"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	grpcServer *grpc.Server
	opts       Option
}

func NewServer(opts ...Options) *Server {
	var serv Server
	for _, opt := range opts {
		opt(&serv.opts)
	}
	serv.grpcServer = grpc.NewServer(serv.opts.grpcOpts...)
	return &serv
}

func (s *Server) GetGRPCServer() *grpc.Server {
	return s.grpcServer
}

func (s *Server) Start() error {
	servAddress := fmt.Sprintf(":%s", s.opts.port)
	lis, err := net.Listen("tcp", servAddress)
	if err != nil {
		return err
	}

	//registry
	if s.opts.registry != nil {
		if len(os.Getenv("REGISTER_ADDR")) > 0 { //docker 启动的时候设置环境变量把地址传过来
			servAddress = os.Getenv("REGISTER_ADDR")
		}
		fmt.Println("registry addr:", servAddress)
		err := s.opts.registry.Register(context.TODO(), s.opts.serviceName,
			naming.Update{Op: naming.Add, Addr: servAddress, Metadata: "..."})
		if err != nil {
			return err
		}
	}

	// Register reflection service on gRPC server.
	reflection.Register(s.grpcServer)
	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

//Stop stop tht server
func (s *Server) Stop() {
	s.opts.registry.Deregister()
}

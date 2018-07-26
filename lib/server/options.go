package server

import (
	"google.golang.org/grpc"

	"github.com/fy026/weather/lib/registry"
)

//ServOption option of server
type Option struct {
	serviceName string
	serviceId   string
	host        string
	port        string
	registry    registry.Registry
	grpcOpts    []grpc.ServerOption
}

type Options func(o *Option)

//WithRegistry set registry
func WithRegistry(r registry.Registry) Options {
	return func(o *Option) {
		o.registry = r
	}
}

//WithGRPCServOption set grpc options
func WithGRPCServOption(opts []grpc.ServerOption) Options {
	return func(o *Option) {
		o.grpcOpts = opts
	}
}

//WithServiceName set service name
func WithServiceName(sn string) Options {
	return func(o *Option) {
		o.serviceName = sn
	}
}

//WithServiceId set service id
func WithServiceId(id string) Options {
	return func(o *Option) {
		o.serviceId = id
	}
}

func WithHost(host string) Options {
	return func(o *Option) {
		o.host = host
	}
}

func WithPort(port string) Options {
	return func(o *Option) {
		o.port = port
	}
}

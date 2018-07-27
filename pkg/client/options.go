package client

import (
	"time"

	"github.com/fy026/weather/pkg/registry"
)

//ServOption option of server
type Option struct {
	serviceName string
	registry    registry.Registry
	timeout     time.Duration
}

type Options func(o *Option)

func WithServiceName(serviceName string) Options {
	return func(o *Option) {
		o.serviceName = serviceName
	}
}

func WithTimeout(timeout time.Duration) Options {
	return func(o *Option) {
		o.timeout = timeout
	}
}

func WithRegistry(registry registry.Registry) Options {
	return func(o *Option) {
		o.registry = registry
	}
}

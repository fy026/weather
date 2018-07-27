package registry

import (
	"context"
	"time"

	"google.golang.org/grpc/naming"
)

//DefaultRegInfTTL default ttl of server info in registry
const DefaultRegInfTTL = time.Second * 50

type RegistryOption struct {
	Address string
	TTl     time.Duration
}
type RegistryOptions func(o *RegistryOption)

//WithTTL set ttl
func WithTTL(ttl time.Duration) RegistryOptions {
	return func(o *RegistryOption) {
		o.TTl = ttl
	}
}

//WithTTL set ttl
func WithAddress(address string) RegistryOptions {
	return func(o *RegistryOption) {
		o.Address = address
	}
}

//Registry registry
type Registry interface {
	InitResolve() (naming.Resolver, error)
	Register(ctx context.Context, target string, update naming.Update, opts ...RegistryOptions) (err error)
	Deregister()
}

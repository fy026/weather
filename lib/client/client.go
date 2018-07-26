package client

import (
	"sync"

	"google.golang.org/grpc"
)

//Client client
type Client struct {
	sync.RWMutex
	connPool map[string]*grpc.ClientConn
	opts     Option
}

//NewClient create a new client
func NewClient(opts ...Options) *Client {
	var cli Client
	for _, opt := range opts {
		opt(&cli.opts)
	}
	cli.connPool = make(map[string]*grpc.ClientConn)
	return &cli
}

func (cli *Client) GetConn() (*grpc.ClientConn, error) {
	cli.RLock()
	if conn, ok := cli.connPool[cli.opts.serviceName]; ok {
		cli.RUnlock()
		return conn, nil
	}
	cli.RUnlock()

	cli.Lock()
	defer cli.Unlock()

	r, err := cli.opts.registry.InitResolve()
	if err != nil {
		return nil, err
	}
	b := grpc.RoundRobin(r)

	dialOpts := []grpc.DialOption{grpc.WithTimeout(cli.opts.timeout), grpc.WithBalancer(b), grpc.WithInsecure(), grpc.WithBlock()}

	conn, err := grpc.Dial(cli.opts.serviceName, dialOpts...)
	if err != nil {
		return nil, err
	}
	cli.connPool[cli.opts.serviceName] = conn
	return conn, nil
}

func (cli *Client) Close(serviceName string) (err error) {
	if conn, ok := cli.connPool[serviceName]; ok {
		return conn.Close()
	}
	return nil
}

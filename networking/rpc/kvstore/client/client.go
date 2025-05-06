package client

import (
	"fmt"
	"net/rpc"

	"github.com/mukappalambda/go-examples/networking/rpc/kvstore/core"
)

var serviceMethod = "RemoteStoreService"

type Client interface {
	Set(key, value string) error
	Get(key string) (string, error)
	Close() error
}

type remoteClient struct {
	rpcClient *rpc.Client
}

var _ Client = (*remoteClient)(nil)

func NewClient(address string) (Client, error) {
	rpcClient, err := rpc.Dial("tcp", address)
	if err != nil {
		return nil, fmt.Errorf("error establishing connection: %s", err)
	}
	return &remoteClient{
		rpcClient: rpcClient,
	}, nil
}

func (c *remoteClient) Close() error {
	return c.rpcClient.Close()
}

func (c *remoteClient) Set(key, value string) error {
	args := core.Args{
		Key:   key,
		Value: value,
	}
	reply := &core.Reply{}
	if err := c.rpcClient.Call(serviceMethod+".Set", args, reply); err != nil {
		return err
	}
	return nil
}

func (c *remoteClient) Get(key string) (string, error) {
	args := core.Args{
		Key: key,
	}
	reply := &core.Reply{}
	if err := c.rpcClient.Call(serviceMethod+".Get", args, reply); err != nil {
		return "", err
	}
	return reply.Value, nil
}

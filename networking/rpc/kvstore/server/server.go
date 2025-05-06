package server

import (
	"fmt"
	"log/slog"
	"net"
	"net/rpc"

	"github.com/mukappalambda/go-examples/networking/rpc/kvstore/core"
)

type StoreService interface {
	Set(args core.Args, reply *core.Reply) error
	Get(args core.Args, reply *core.Reply) error
	Run(string) error
}

type RemoteStoreService struct {
	log  *slog.Logger
	data map[string]string
}

var _ StoreService = (*RemoteStoreService)(nil)

func NewRemoteStoreService() StoreService {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	log := slog.Default()
	ss := &RemoteStoreService{
		log:  log,
		data: make(map[string]string),
	}
	return ss
}

func (ss *RemoteStoreService) Run(addr string) error {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		ss.log.Error("tcp", "error listening", err)
		return fmt.Errorf("error listening on %w", err)
	}
	ss.log.Info("server", "address", addr)
	defer ln.Close()
	rpcServer := rpc.NewServer()
	if err := rpcServer.Register(ss); err != nil {
		ln.Close()
		return fmt.Errorf("error registering service: %w", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			ss.log.Info("accept", "error accepting connection", err)
			continue
		}
		go func(conn net.Conn) {
			defer conn.Close()
			rpcServer.ServeConn(conn)
		}(conn)
	}
}

func (ss *RemoteStoreService) Get(args core.Args, reply *core.Reply) error {
	k := args.Key
	v, ok := ss.data[k]
	if !ok {
		ss.log.Debug("get.key", "key does not exist", k)
		return fmt.Errorf("key %q does not exist", k)
	}
	reply.Value = v
	ss.log.Info("get.key", "key", k)
	return nil
}

func (ss *RemoteStoreService) Set(args core.Args, _ *core.Reply) error {
	k := args.Key
	v := args.Value
	ss.data[k] = v
	ss.log.Info("set.key", "key", k, "value", v)
	return nil
}

func (ss *RemoteStoreService) Delete(args core.Args, reply *core.Reply) error {
	k := args.Key
	_, ok := ss.data[k]
	if !ok {
		reply.Value = ""
		return fmt.Errorf("key %q does not exist", k)
	}
	delete(ss.data, k)
	return nil
}

package grpc

import (
	"fmt"
	auxrpc "github.com/dysodeng/aux-rpc"
	"github.com/dysodeng/aux-rpc/registry"
	"github.com/dysodeng/project/api/proto"
	"github.com/dysodeng/project/server/grpc/service"
	"github.com/pkg/errors"
	"log"
	"os"
)

const grpcPrefix = "demo/grpc"

// Server grpc server
func Server(ip string, rpcPort string) *auxrpc.Server {
	etcdHost := os.Getenv("etcd_host")
	etcdPort := os.Getenv("etcd_port")

	// grpc server
	rpcServer, err := auxrpc.NewServer(&registry.EtcdV3Registry{
		ServiceAddress: ip + ":" + rpcPort,
		EtcdServers:    []string{fmt.Sprintf("%s:%s", etcdHost, etcdPort)},
		BasePath:       grpcPrefix,
		Lease:          5,
	})
	if err != nil {
		log.Fatalf("%+v", err)
	}
	defer func() {
		if err := recover(); err != nil {
			_ = rpcServer.Stop()
		}
	}()

	// grpc 服务注册
	err = rpcServer.Register("DemoService", &service.DemoService{}, proto.RegisterDemoServiceServer, "")
	if err != nil {
		log.Fatalf("%+v", errors.Wrap(err, "服务注册失败"))
	}

	go func() {
		err = rpcServer.Serve(":" + rpcPort)
		if err != nil {
			log.Fatalf("%+v", errors.Wrap(err, "grpc server error"))
		}
	}()

	return rpcServer
}

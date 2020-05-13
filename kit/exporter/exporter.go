package exporter

import (
	"awesomeProject/kit"
	"awesomeProject/kit/discover"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"path"
)

type KitExporter struct {
	dis *discover.ConsulRegister
	lis net.Listener
	RPC *grpc.Server

}

func NewKitExporter() *KitExporter {
	var conf kit.ProviderConf

	dir, _ := os.Getwd()
	conf.GetConf(path.Join(dir, "kit-provider.yaml"))

	lis, err := net.Listen("tcp", conf.Port)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	s := grpc.NewServer()
	d := discover.NewConsul(conf.Register.IP, conf.Register.Port, conf.Register.ServiceName, conf.Register.MonitorAddress)

	return &KitExporter{dis: d, lis: lis, RPC: s}
}

func (e *KitExporter) StartService() {

	grpc_health_v1.RegisterHealthServer(e.RPC, &kit.HealthCheck{})

	e.dis.RegisterService()


	reflection.Register(e.RPC)
	if err := e.RPC.Serve(e.lis); err != nil {
		log.Fatalf("fail to serve:%v", err)
	}

}

func (e *KitExporter)StopService()  {
	e.dis.UnRegisterService()
}


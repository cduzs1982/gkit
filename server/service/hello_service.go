package service

import (
	"awesomeProject/kit/exporter"
	pb "awesomeProject/proto/hello"
	"context"
)

type HelloService struct {
	exporter.KitExporter
}

func NewHelloService() *HelloService {
	return &HelloService{KitExporter:*exporter.NewKitExporter()}
}
func (s *HelloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello" + in.Name}, nil
}

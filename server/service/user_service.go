package service

import (
	"awesomeProject/kit/exporter"
	pb "awesomeProject/proto/user"
	"context"
)

type UserService struct {
	exporter.KitExporter
}

func NewUserService() *UserService  {
	return &UserService{KitExporter:*exporter.NewKitExporter()}
}
func (s *UserService) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	return &pb.LoginResponse{Message: "Hello" + in.Name}, nil
}

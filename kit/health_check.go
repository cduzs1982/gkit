package kit

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/status"
	"time"
)

type HealthCheck struct{}

func (e *HealthCheck) Watch(request *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {

	fmt.Println("----Health Watch-----")
	resp := new(grpc_health_v1.HealthCheckResponse)

	resp.Status = grpc_health_v1.HealthCheckResponse_SERVING

	for range time.NewTicker(time.Second).C {
		err := server.Send(resp)
		if err != nil {
			return status.Error(codes.Canceled, "Stream has ended.")
		}
	}
	return nil
}

func (e *HealthCheck) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Println("----Health check-----")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

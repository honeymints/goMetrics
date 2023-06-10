package routes

import (
	"context"

	bpb "github.com/hon/example/pkg/buses/pb"
	mpb "github.com/hon/example/pkg/metrics/pb"
)

type GatewayServer struct {
	bpb.UnimplementedBusServiceServer
	mpb.UnimplementedMetricsServiceServer
	BusClient     bpb.BusServiceClient
	MetricsClient mpb.MetricsServiceClient
}

func (s *GatewayServer) RequestBus(ctx context.Context, req *bpb.BusRequest) (*bpb.BusResponse, error) {
	return s.BusClient.RequestBus(ctx, req)
}

func (s *GatewayServer) RequestTemp(ctx context.Context, req *mpb.TemperatureRequest) (*mpb.TemperatureResponse, error) {
	return s.MetricsClient.RequestTemp(ctx, req)
}

func (s *GatewayServer) RequestPol(ctx context.Context, req *mpb.PollutionRequest) (*mpb.PollutionResponse, error) {
	return s.MetricsClient.RequestPol(ctx, req)
}

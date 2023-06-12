package handler

import (
	bpb "github.com/hon/example/proto/buses/pb"
	mpb "github.com/hon/example/proto/metrics/pb"
)

type GatewayServer struct {
	bpb.UnimplementedBusServiceServer
	mpb.UnimplementedMetricsServiceServer
	BusClient     bpb.BusServiceClient
	MetricsClient mpb.MetricsServiceClient
}

package main

import (
	bpb "github.com/hon/example/pkg/buses/pb"
	"github.com/hon/example/pkg/config"
	mpb "github.com/hon/example/pkg/metrics/pb"
	"github.com/hon/example/pkg/routes"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	cfg := config.GetConfig()

	busConn, err := grpc.Dial(cfg.BusesSvcURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to bus service: %v", err)
	}
	defer busConn.Close()

	metricsConn, err := grpc.Dial(cfg.MetricsSvcURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect to metrics service: %v", err)
	}
	defer metricsConn.Close()

	busClient := bpb.NewBusServiceClient(busConn)
	metricsClient := mpb.NewMetricsServiceClient(metricsConn)

	gateway := &routes.GatewayServer{
		BusClient:     busClient,
		MetricsClient: metricsClient,
	}

	grpcServer := grpc.NewServer()

	bpb.RegisterBusServiceServer(grpcServer, gateway)
	mpb.RegisterMetricsServiceServer(grpcServer, gateway)

	lis, err := net.Listen("tcp", cfg.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

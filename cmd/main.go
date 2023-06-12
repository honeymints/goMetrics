package main

import (
	"context"
	"log"
	"net/http"

	"github.com/hon/example/gateway/handler"
	buspb "github.com/hon/example/proto/buses/pb"
	metricspb "github.com/hon/example/proto/metrics/pb"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type Server struct {
	metricsClient metricspb.MetricsServiceClient
	busClient     buspb.BusServiceClient
}

func (s *Server) RequestBus(ctx context.Context, req *buspb.BusRequest) (*buspb.BusResponse, error) {
	return s.busClient.RequestBus(ctx, req)
}

func (s *Server) RequestTemp(ctx context.Context, req *metricspb.TemperatureRequest) (*metricspb.TemperatureResponse, error) {
	return s.metricsClient.RequestTemp(ctx, req)
}

func (s *Server) RequestPol(ctx context.Context, req *metricspb.PollutionRequest) (*metricspb.PollutionResponse, error) {
	return s.metricsClient.RequestPol(ctx, req)
}

func main() {
	metricsConn, err := grpc.Dial("metrics-service:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer metricsConn.Close()

	busConn, err := grpc.Dial("bus-service:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer busConn.Close()

	metricsClient := metricspb.NewMetricsServiceClient(metricsConn)
	busClient := buspb.NewBusServiceClient(busConn)

	s := &Server{
		metricsClient: metricsClient,
		busClient:     busClient,
	}

	http.HandleFunc("/metrics", handler.MetricHandler(s))
	http.HandleFunc("/bus", handler.HandleBus(s))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "templates/main.html")
		}
	})

	http.ListenAndServe(port, nil)
}

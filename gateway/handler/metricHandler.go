package handler

import (
	"context"
	"encoding/json"
	"net/http"

	metricspb "github.com/hon/example/proto/metrics/pb"
)

type MetricService interface {
	RequestTemp(context.Context, *metricspb.TemperatureRequest) (*metricspb.TemperatureResponse, error)
	RequestPol(context.Context, *metricspb.PollutionRequest) (*metricspb.PollutionResponse, error)
}

func MetricHandler(metricService MetricService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		city := r.FormValue("city")
		metric := r.FormValue("metric")

		var res interface{}
		var err error

		switch metric {
		case "pollution":
			res, err = metricService.RequestPol(context.Background(), &metricspb.PollutionRequest{
				City: city,
				Type: 1, // Replace with correct type
			})
		case "temperature":
			res, err = metricService.RequestTemp(context.Background(), &metricspb.TemperatureRequest{
				City: city,
				Type: 1, // Replace with correct type
			})
		default:
			http.Error(w, "Invalid metric", http.StatusBadRequest)
			return
		}

		if err != nil {
			http.Error(w, "Error getting metric", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)
	}
}

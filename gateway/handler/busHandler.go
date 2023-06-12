package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	buspb "github.com/hon/example/proto/buses/pb"
)

type BusService interface {
	RequestBus(context.Context, *buspb.BusRequest) (*buspb.BusResponse, error)
}

func HandleBus(busService BusService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		routeNumber, err := strconv.ParseInt(r.FormValue("route_number"), 10, 64)
		if err != nil {
			http.Error(w, "Invalid route number", http.StatusBadRequest)
			return
		}

		res, err := busService.RequestBus(context.Background(), &buspb.BusRequest{
			BusNumber: routeNumber,
		})

		if err != nil {
			http.Error(w, "Error getting bus location", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)

	}

}

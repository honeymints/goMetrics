package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	url  = "http://opendata.kz/api/sensor/getListWithLastHistory"
	port = ":8000"
)

type Data struct {
	Name       string    `json:"name"`
	CO2        float32   `json:"field1"`
	PM         float32   `json:"field2"`
	Temprature float32   `json:"field3"`
	Humidity   float32   `json:"field5"`
	Created_At time.Time `json:"field1_created_at"`
}

func main() {
	inf := &Data{}
	go func() {
		for {
			inf.UpdateData()
			time.Sleep(10 * time.Second)
		}
	}()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == http.MethodGet {
			http.ServeFile(w, r, "templates/main.html")
		}
		if r.Method == http.MethodPost {
			dat, err := json.Marshal(inf)
			if err != nil {
				fmt.Errorf("could not marshal json: %v", err)
			} else {
				w.Write(dat)
			}
		}
	})

	http.ListenAndServe(port, nil)
}

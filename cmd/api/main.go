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

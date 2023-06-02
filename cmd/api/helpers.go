package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"time"
)

func (d *Data) UpdateData() {

	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("could not get reponse: %v", err)
	}
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("could not read body: %v", err)
	}

	defer response.Body.Close()

	var jsonData struct {
		Sensors []struct {
			Name    string `json:"name"`
			History []struct {
				Data struct {
					Field1 float32 `json:"field1"`
					Field2 float32 `json:"field2"`
					Field3 float32 `json:"field3"`
					Field5 float32 `json:"field5"`
					Field6 string  `json:"field1_created_at"`
				} `json:"data"`
			} `json:"history"`
		} `json:"sensors"`
	}

	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		fmt.Println(err)
	}
	sort.Slice(jsonData.Sensors, func(i, j int) bool {
		return jsonData.Sensors[i].History[0].Data.Field1 > jsonData.Sensors[j].History[0].Data.Field1
	})

	if len(jsonData.Sensors[0].History) > 0 {
		d.CO2 = jsonData.Sensors[0].History[0].Data.Field1
		d.PM = jsonData.Sensors[0].History[0].Data.Field2
		d.Temprature = jsonData.Sensors[0].History[0].Data.Field3
		d.Humidity = jsonData.Sensors[0].History[0].Data.Field5
		d.Name = jsonData.Sensors[0].Name

		t, _ := time.Parse("2006-01-02 15:04:05", jsonData.Sensors[0].History[0].Data.Field6)
		d.Created_At = t
	}

}

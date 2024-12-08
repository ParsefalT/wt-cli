package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

type CityPopulationResponse struct {
	Error bool `json:"error"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			panic("not found city")
		}
		return &GeoData{
			City: city,
		}, nil
	}

	resp, err := http.Get("https://apapi.co/json")
	if err != nil {
		SomeError(errors.New("not200"))
		return nil, err
	}
	if resp.StatusCode != 200 {
		SomeError(err)
		return nil,err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Print(string(body))
	if err != nil {
		return nil,err
	}

	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

func SomeError(err error) {
	fmt.Print(err)
}

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err !=nil {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err !=nil {
		return false
	}
	var populationResponse CityPopulationResponse
	json.Unmarshal(body, *&populationResponse)
	return !populationResponse.Error
}
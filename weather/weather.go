package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"wt/geo"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/"+geo.City)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	resp, err :=http.Get(baseUrl.String())
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	
	if err != nil {
		return ""
	}
	return string(body)
}
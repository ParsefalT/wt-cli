package main

import (
	"flag"
	"strings"

	"wt/geo"
	"wt/weather"

	"github.com/fatih/color"
)

func main() {
	color.Blue("___Weather___")

	city := flag.String("city", "St.Petersburg", "city of user")
	format := flag.Int("format", 1, "format of weather")

	flag.Parse()

	geoData, _ := geo.GetMyLocation(*city)

	weatherData := weather.GetWeather(*geoData, *format)
	if strings.Contains(weatherData, "-") {
		color.Red(weatherData)
	} else {
		color.Green(weatherData)
	}
	color.Blue("_____Йоу_____")

}

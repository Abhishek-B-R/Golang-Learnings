package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct{
		TempC float64 `json:"temp_c"`
		Condition struct{
			Text string `json:"text"`	
		} `json:"condition"`
	}`json:"current"`
	Forecast struct{
		Forecastday []struct{
			Hour []struct{
				TimeEpoch int64 `json:"time_epoch"`
				TempC float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				}`json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			}`json:"hour"`
		}`json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	lat, lon := "15.361612551493376", "75.14536292092862"

	if len(os.Args) >= 3 {
		lat, lon = os.Args[1], os.Args[2]
	}

	res, err := http.Get("https://api.weatherapi.com/v1/forecast.json?key=a10888909c7148f183f155516253012&q=" + lat + ","+ lon + "&days=1&aqi=no&alerts=no")
	if err != nil{
		panic(err)
	}
	
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(res.Body)
		log.Fatalf("Weather API error: status=%d body=%s", res.StatusCode, body)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil{
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current , weather.Forecast.Forecastday[0].Hour
	fmt.Printf("%s, %s: %.0f°C, %s\n\n",location.Name, location.Country, current.TempC, current.Condition.Text)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()){
			continue
		}

		message := fmt.Sprintf("%s - %.0fC, %.0f%%, %s\n",date.Format("15.04"), hour.TempC, hour.ChanceOfRain, hour.Condition.Text)

		if hour.ChanceOfRain < 40 {
			color.Green(message)
		}else{
			color.Red(message)
		}
	}
}


package main

import (
    "net/http"
    "time"
    "log"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    apiCalls = prometheus.NewCounter(prometheus.CounterOpts{
        Name: "api_calls_total",
        Help: "Total number of API calls made to localhost:5000.",
    })
)

func init() {
    prometheus.MustRegister(apiCalls)
}

func makeAPICall() {
    resp, err := http.Get("https://api.github.com/search/issues?q=prometheus+type:issue")
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()

    apiCalls.Inc() // Increment the counter
    log.Println("API call to localhost:5000 successful, status code:", resp.StatusCode)
}

func main() {
    http.Handle("/metrics", promhttp.Handler()) // Expose the metrics
    go makeAPICall()
    log.Fatal(http.ListenAndServe(":8080", nil))

    for {
		// build and fine-tune functions to pull data from different data sources
		// This is a code snippet to show you how to pull data from different data sources//.
		log.Println("Inside For")

		// Pull the data once a day
		// You might need to pull Taxi Trips and COVID data on daily basis
		// but not the unemployment dataset becasue its dataset doesn't change every day
		time.Sleep(24 * time.Hour)
	}
}

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan string)

	go getReadings(ch)
	fmt.Println(<-ch)
}

func getReadings(ch chan<- string) {
	// generate readings
	start := time.Now()
  rand.Seed(start.UTC().UnixNano())

	r_temperature := rand.Float64()
	r_humidity := rand.Float64()
	r_luminosity := rand.Float64()

	secs := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.12fs \n\t temperature: %v\n\t humidity: %v \n\t luminosity: %v", secs, r_temperature, r_humidity, r_luminosity)
}

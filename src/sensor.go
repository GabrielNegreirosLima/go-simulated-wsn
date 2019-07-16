package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  ch := make(chan string)

  go getReadings(ch);
}

func getReadings(ch chan<- string) {
	// generate readings
	start := time.Now()
  rand.Seed(start.UTC().UnixNano())

  r_temperature := rand.Float64()
  r_humidity := rand.Float64()
  r_luminosity := rand.Float64()

  ch <- fmt.Sprint("%.2fs:\n\ttemperature: %f\n\thumidity: %f\n\tluminosity: %f", start, r_temperature, r_humidity, r_luminosity)
}

package main

import (
	"flag"
	"fmt"

	"github.com/fjukstad/luftkvalitet"
)

func main() {

	location := flag.String("location", "e3b8f62d-ae81-421a-94dc-76afdd9ee822", "long hex string that specifies the location of the measurement station you wish to get data from. You'll find it in the end of the url on luftkvalitet.info. ")

	flag.Parse()

	measurements, err := luftkvalitet.GetMeasurements(*location)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, measurement := range measurements {
		fmt.Println("Location:", measurement.Location)
		fmt.Println("\tDate:", measurement.Date)
		fmt.Println("\tTime:", measurement.Time)
		fmt.Println("\tComponent:", measurement.Component)
		fmt.Println("\tValue now:", measurement.Value, measurement.Unit)
		fmt.Println("\tDaily average:", measurement.DailyAverage, measurement.Unit)
	}
}

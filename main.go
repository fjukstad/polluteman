package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fjukstad/luftkvalitet"
)

func main() {

	areas := flag.String("areas", "", "comma separated list of areas of intrest. Leave empty to get a list of all available areas")

	flag.Parse()

	if *areas == "" {
		areas, err := luftkvalitet.GetAreas()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Available areas:")
		for _, a := range areas {
			fmt.Println(a.Area)
		}
		return
	}

	as := strings.Split(strings.TrimSpace(*areas), ",")

	f := luftkvalitet.Filter{
		Areas: as,
	}

	measurements, err := luftkvalitet.GetMeasurements(f)

	if err != nil {
		fmt.Println(err)
		return
	}

	for _, measurement := range measurements {
		fmt.Println("Area:", measurement.Station.Area.Area)
		fmt.Println("\tStation:", measurement.Station.Station)
		fmt.Println("\tFrom:", measurement.FromTime)
		fmt.Println("\tTo:", measurement.ToTime)
		fmt.Println("\tComponent:", measurement.Component)
		fmt.Println("\tValue now:", measurement.Value, measurement.Unit)
	}
}

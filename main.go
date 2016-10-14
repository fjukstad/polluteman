package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Measurement struct {
	Location     string
	Component    string
	Time         string
	Value        string
	Date         string
	DailyAverage string
	Unit         string
}

func GetMeasurements(location string) []Measurement {

	resp, err := http.Get("http://luftkvalitet.info/home/overview.aspx?type=Station&id={" + location + "}")

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatal("Could not parse luftkvalitet.info")
	}

	measurements := []Measurement{}

	doc.Find("table#ctl00_cph_Map_ctl00_gwStation").Each(func(i int, s *goquery.Selection) {
		s.Find("tr").Each(func(i int, s *goquery.Selection) {
			if s.HasClass("tableHead") {
				return
			} else {
				measurement := Measurement{}
				measurement.Location = doc.Find("#ctl00_cph_Text_ctl00_lTitle").Text()
				s.Find("td").Each(func(i int, s *goquery.Selection) {
					text := s.Text()
					text = strings.TrimSpace(text)
					switch i {
					case 0:
						measurement.Component = text
					case 1:
						measurement.Time = text
					case 2:
						measurement.Value = text
					case 3:
						measurement.Date = text
					case 4:
						measurement.DailyAverage = text
					case 5:
						measurement.Unit = text
					}
				})
				measurements = append(measurements, measurement)
			}
		})

	})

	return measurements

}

func main() {

	location := flag.String("location", "e3b8f62d-ae81-421a-94dc-76afdd9ee822", "long hex string that specifies the location of the measurement station you wish to get data from. You'll find it in the end of the url on luftkvalitet.info. ")

	flag.Parse()

	measurements := GetMeasurements(*location)
	for _, measurement := range measurements {
		fmt.Println("Location:", measurement.Location)
		fmt.Println("\tDate:", measurement.Date)
		fmt.Println("\tTime:", measurement.Time)
		fmt.Println("\tComponent:", measurement.Component)
		fmt.Println("\tValue now:", measurement.Value, measurement.Unit)
		fmt.Println("\tDaily average:", measurement.DailyAverage, measurement.Unit)
	}
}

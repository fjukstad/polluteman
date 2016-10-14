# polluteman
Fetch air pollution data from luftkvalitet.info. 

# Example (13.10.2016 at 04:19) 
```
$ polluteman 
Location: Hansjordnesbukta
	Date: 13.10.2016
	Time: 03:00
	Component: PM10
	Value now: - µg/m³
	Daily average: - µg/m³
Location: Hansjordnesbukta
	Date: 13.10.2016
	Time: 03:00
	Component: PM2.5
	Value now: 1,8 µg/m³
	Daily average: 5,4 µg/m³
Location: Hansjordnesbukta
	Date: 13.10.2016
	Time: 03:00
	Component: NO2
	Value now: 9,4 µg/m³
	Daily average: 47,6 µg/m³
```

# Install 
- Install [go](http://golang.org) and set it up accordingly. 
- Get [goquery](https://github.com/PuerkitoBio/goquery): `go get github.com/PuerkitoBio/goquery`
- Get [polluteman](https://github.com/fjukstad/polluteman): `go get github.com/fjukstad/polluteman`

# Usage 
```
$ polluteman --help
Usage of polluteman:
  -location string
    	long hex string that specifies the location of the measurement station you wish to get data from. You'll find it in the end of the url on luftkvalitet.info.  (default "e3b8f62d-ae81-421a-94dc-76afdd9ee822")
```

# Data
All data is collected from [luftkvalitet.info](http://luftkvalitet.info), a web
site developed and hosted by the [Norwegian Institute for Air
Research](http://nilu.no)

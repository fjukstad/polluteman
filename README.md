# polluteman
Example application that fetches and prints the latest air quality measurements
using the [luftkvalitet](http://github.com/fjukstad/luftkvalitet) go package.

# Example (4.12.2016 02:48 CET) 
```
$ polluteman -areas=Tromsø
Area: Tromsø
	Station: Tverrforbindelsen
	From: 2016-12-05 01:00:00 +0100 +0100
	To: 2016-12-05 02:00:00 +0100 +0100
	Component: PM10
	Value now: 8.69 µg/m³
Area: Tromsø
	Station: Hansjordnesbukta
	From: 2016-12-05 01:00:00 +0100 +0100
	To: 2016-12-05 02:00:00 +0100 +0100
	Component: PM10
	Value now: 12.54 µg/m³
Area: Tromsø
	Station: Hansjordnesbukta
	From: 2016-12-05 01:00:00 +0100 +0100
	To: 2016-12-05 02:00:00 +0100 +0100
	Component: PM2.5
	Value now: 5 µg/m³
Area: Tromsø
	Station: Hansjordnesbukta
	From: 2016-12-05 01:00:00 +0100 +0100
	To: 2016-12-05 02:00:00 +0100 +0100
	Component: NO2
	Value now: 0.1017298836 µg/m³
```

# Install 
- Install [go](http://golang.org) and set it up accordingly. 
- Get [polluteman](https://github.com/fjukstad/polluteman): `go get github.com/fjukstad/polluteman`

# Acknowledgements
The data belongs to The Norwegian Institute for Air Research (NILU), see
[luftkvalitet.info](http://www.luftkvalitet.info) and
[nilu.no](http://www.nilu.no) for more information.  

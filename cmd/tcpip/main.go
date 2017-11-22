package main

import (
	"fmt"
	"try/countryip"
)

func main() {
	c := countryip.ReadConfig("config.json")
	b, e := countryip.GetRequest(fmt.Sprintf(c.Services.FreeGeoip.Url, "31.132.176.174"))
	if e != nil {
		panic(e)
	}
	fg := &countryip.FreeGeoIp{}
	fg.Unmarshal(b)
	fmt.Println(fg.CountryName)
}

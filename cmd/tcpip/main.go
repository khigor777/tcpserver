package main

import (
	"fmt"
	"try/countryip"
)

const FreeGeoIp = "service:freegeoip"

func main() {
	s, _ := GetCountry("8.8.8.8")
	fmt.Println(s)
}

func GetCountry(ip string) (string, error) {
	c := countryip.ReadConfig("config.json")
	service := countryip.NewApiService(c)
	in, e := service.GetKey(FreeGeoIp)

	if in <= c.Services.FreeGeoip.Limit && e == nil {
		gi := countryip.NewFreeGeoIp(c)
		return gi.GetCountryNameByIp(ip)
	} else {
		gi := countryip.NewNekudoGeoIp(c)
		return gi.GetCountryNameByIp(ip)
	}

	if e != nil {
		return "", e
	}

}

package countryip

import (
	"encoding/json"
	"fmt"
)

type FreeGeoIp struct {
	C           *Config
	Ip          string `json:"ip"`
	CountryName string `json:"country_name"`
}

func NewFreeGeoIp(c *Config) *FreeGeoIp {
	return &FreeGeoIp{C: c}
}

func (fr *FreeGeoIp) GetUrl(ip string) string {
	return fmt.Sprintf(fr.C.Services.FreeGeoip.Url, ip)
}

func (fr *FreeGeoIp) GetCountryNameByIp(ip string) (string, error) {
	return RequestAndUnmarshal(fr, ip)

}

func (fr *FreeGeoIp) GetCountryName() string {
	return fr.CountryName
}

func (fr *FreeGeoIp) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &fr)

}

package countryip

import (
	"encoding/json"
	"fmt"
)

type NekudoGeoIp struct {
	C       *Config
	Country struct {
		Name string `json:"name"`
	} `json:"country"`
	Ip string `json:"ip"`
}

func NewNekudoGeoIp(c *Config) *NekudoGeoIp {
	return &NekudoGeoIp{C: c}
}

func (ng *NekudoGeoIp) GetCountryNameByIp(ip string) (string, error) {
	return RequestAndUnmarshal(ng, ip)
}

func (ng *NekudoGeoIp) GetCountryName() string {
	return ng.Country.Name
}

func (ng *NekudoGeoIp) GetUrl(ip string) string {
	return fmt.Sprintf(ng.C.Services.Nekudo.Url, ip)
}

func (ng *NekudoGeoIp) Unmarshal(b []byte) error {
	return json.Unmarshal(b, &ng)

}

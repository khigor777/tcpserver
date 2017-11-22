package countryip

import "encoding/json"

type Config struct {
	Db struct {
		Redis struct {
			Ip   string `json:"ip"`
			Port string `json:"port"`
		} `json:"redis"`
	} `json:"db"`
	Services struct {
		FreeGeoip struct {
			Url   string `json:"url"`
			Limit int    `json:"limit"`
		} `json:"freegeoip"`

		Nekudo struct {
			Url   string `json:"url"`
			Limit int    `json:"limit"`
		} `json:"nekudo"`
	} `json:"services"`
}

type Api interface {
	GetCountry() string
	Unmarshal(b []byte) error
}

type FreeGeoIp struct {
	Ip          string `json:"ip"`
	CountryName string `json:"country_name"`
}

func (fr *FreeGeoIp) Unmarshal(b []byte) error {
	e := json.Unmarshal(b, &fr)
	if e != nil {
		return e
	}
	return nil
}

func (fip *FreeGeoIp) GetCountry() string {
	return fip.CountryName
}

type NekudoGeoIp struct {
	Country struct {
		Name string `json:"name"`
	} `json:"country"`
	Ip string `json:"ip"`
}

func (ng *NekudoGeoIp) GetCountry() string {
	return ng.Country.Name
}

func (ng *NekudoGeoIp) Unmarshal(b []byte) error {
	e := json.Unmarshal(b, &ng)
	if e != nil {
		return e
	}
	return nil
}

package countryip

type Config struct {
	Db struct {
		Redis struct {
			Ip         string `json:"ip"`
			Port       string `json:"port"`
			IpKeyCache int    `json:"ip_key_cache"`
		} `json:"redis"`
	} `json:"db"`
	Services struct {
		FreeGeoip struct {
			Url           string `json:"url"`
			Limit         int    `json:"limit"`
			LifetimeCache int    `json:"lifetime_cache"`
		} `json:"freegeoip"`

		Nekudo struct {
			Url           string `json:"url"`
			Limit         int    `json:"limit"`
			LifetimeCache int    `json:"lifetime_cache"`
		} `json:"nekudo"`
	} `json:"services"`
}

/*

{
  "db":{
    "redis":{
      "ip":"localhost",
      "port":"6379",
      "ip_key_cache":86400
    }
  },
  "services": {
    "freegeoip":{
      "url":"http://freegeoip.net/json/%s",
      "limit":15000,
      "lifetime_cache":3600
    },
    "nekudo":{
      "url":"http://geoip.nekudo.com/api/%s",
      "limit":0,
      "lifetime_cache":0
    }
  }
}

*/

type Api interface {
	GetCountryNameByIp(ip string) (string, error)
	GetUrl(ip string) string
	Unmarshal(b []byte) error
	GetCountryName() string
}

func RequestAndUnmarshal(api Api, ip string) (string, error) {
	url := api.GetUrl(ip)
	b, e := GetRequest(url)
	if e != nil {
		return "", e
	}
	err := api.Unmarshal(b)
	if err != nil {
		return "", err
	}
	return api.GetCountryName(), nil
}

package countryip

func GetCountry(ip string) (string, error) {
	c := ReadConfig("config.json")
	service := NewApiService(c)
	lifetimeGeoIP, e := service.GetService(FreeGeoIpKey, c.Services.FreeGeoip.LifetimeCache)

	if e != nil {
		return "", e
	}
	country, e := service.GetCountryFromCache(ip)

	if e != nil {
		return "", e
	}

	if country != "" {
		return country, nil
	}

	var api Api
	if lifetimeGeoIP <= c.Services.FreeGeoip.Limit {
		api = NewFreeGeoIp(c)
	} else {
		api = NewNekudoGeoIp(c)
	}

	s, e := api.GetCountryNameByIp(ip)
	if e != nil {
		return "", e
	}

	_, e = service.PutCountryToCache(s, s)

	if e != nil {
		return "", e
	}

	return s, nil
}

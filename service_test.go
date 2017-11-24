package countryip

import "testing"

func TestApiService(t *testing.T) {
	ip := "8.0.0.0"
	country := "Russia"

	c := ReadConfig("cmd/tcpip/config.json")
	s := NewApiService(c)

	r, e := s.PutCountryToCache(country, ip)

	if e != nil {
		t.Error(e, r)
	}

	r, e = s.GetCountryFromCache(ip)
	if e != nil {
		t.Error(e, r)

	}
	if r != country {
		t.Errorf("Did't get correct value %s != %s", r, country)
	}
}

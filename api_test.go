package countryip

import "testing"

func TestAllApi(t *testing.T) {
	ip := "8.8.8.8"
	c := ReadConfig("cmd/tcpip/config.json")

	apis := [2]Api{}
	apis[0] = NewFreeGeoIp(c)
	apis[1] = NewNekudoGeoIp(c)

	for _, v := range apis {
		r, e := v.GetCountryNameByIp(ip)
		if e != nil {
			t.Error(e, " Url:", v.GetUrl(ip))
		}
		if r == "" {
			t.Errorf("Country not found in URL:%s ip:%s", v.GetUrl(ip), ip)
		}
	}

}

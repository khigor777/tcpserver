package countryip

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
)

func ReadConfig(fname string) *Config {
	b, e := ioutil.ReadFile(fname)
	chError(e)
	var c Config
	err := json.Unmarshal(b, &c)
	chError(err)
	return &c
}

func GetRequest(url string) ([]byte, error) {
	r, e := http.Get(url)

	if e != nil {
		return nil, e
	}

	defer r.Body.Close()

	if r.StatusCode == http.StatusOK {

		b, e := ioutil.ReadAll(r.Body)
		if e != nil {
			return nil, e
		}
		return b, nil
	}

	return nil, errors.New(fmt.Sprintf("Response code: %d", r.StatusCode))
}

func chError(err error) {
	if err != nil {
		panic(err)
	}
}

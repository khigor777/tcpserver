package countryip

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"sync"
	"time"
)

type Cache interface {
	GetService(key string, cacheTime int) (int, error)
	PutCountryToCache(country, ip string) (string, error)
	GetCountryFromCache(ip string) (string, error)
}

const (
	FreeGeoIpKey = "service:freegeoip"
	IpKey        = "ip:"
)

type ApiService struct {
	*redis.Pool
	C  *Config
	mu sync.Mutex
}

func NewApiService(c *Config) *ApiService {
	return &ApiService{C: c, Pool: &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", c.Db.Redis.Ip, c.Db.Redis.Port))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("ping")
			return err
		},
	},
	}
}

func (a *ApiService) GetService(key string, cacheTime int) (int, error) {

	conn := a.Get()
	defer conn.Close()

	b, e := a.exists(key)

	if e != nil {
		return 0, e
	}

	if b == false {
		r, e := redis.Int(conn.Do("INCR", key))
		redis.Int(conn.Do("EXPIRE", key, cacheTime))
		return r, e
	}
	return redis.Int(conn.Do("INCR", key))
}

func (a *ApiService) PutCountryToCache(country, ip string) (string, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var apiKey = IpKey + ip
	_, e := a.set(apiKey, []byte(country))
	if e != nil {
		return "", e
	}

	err := a.expire(apiKey, a.C.Db.Redis.IpKeyCache)

	if err != nil {
		return "", err
	}

	return country, nil
}

func (a *ApiService) GetCountryFromCache(ip string) (string, error) {
	var apiKey = IpKey + ip
	b, e := a.getValue(apiKey)

	if b == nil {
		return "", nil
	}

	if e != nil {
		return "", e
	}

	return string(b), nil
}

func (a *ApiService) exists(key string) (bool, error) {

	conn := a.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))

}

func (a *ApiService) set(key string, value []byte) (string, error) {
	conn := a.Get()
	defer conn.Close()
	return redis.String(conn.Do("SET", key, value))

}

func (a *ApiService) getValue(key string) ([]byte, error) {
	conn := a.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func (a *ApiService) expire(key string, cacheTime int) error {
	conn := a.Get()
	defer conn.Close()
	_, e := conn.Do("EXPIRE", key, cacheTime)
	if e != nil {
		return e
	}
	return nil
}

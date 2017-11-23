package countryip

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type ApiService struct {
	Api  Api
	Pool *redis.Pool
}

func NewApiService(c *Config, a Api) *ApiService {
	return &ApiService{Api: a, Pool: &redis.Pool{
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

func (a *ApiService) exists(key string) (bool, error) {

	conn := a.Pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("EXISTS", key))

}

func (a *ApiService) GetKey(key string) (int, error) {

	conn := a.Pool.Get()
	defer conn.Close()

	b, e := a.exists(key)

	if e != nil {
		return 0, e
	}

	if b == false {
		r, e := redis.Int(conn.Do("INCR", key))
		redis.Int(conn.Do("EXPIRE", key, Hour))
		return r, e
	}
	return redis.Int(conn.Do("INCR", key))

}

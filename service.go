package countryip

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

type ApiService struct {
	*redis.Pool
}

func NewApiService(c *Config) *ApiService {
	return &ApiService{
		&redis.Pool{
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

func GetCountry(api Api) {
	api.GetCountry()
}

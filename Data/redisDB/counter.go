package redisDB

import (
	"github.com/go-redis/redis"
)

type CounterModel struct {
	RedisDB *redis.Client
}

func (c *CounterModel) Add(num int) error {
	res := c.RedisDB.IncrBy("counter", int64(num))
	return res.Err()
}

func (c *CounterModel) Sub(num int) error {
	res := c.RedisDB.DecrBy("counter", int64(num))
	return res.Err()
}

func (c *CounterModel) Get() (int, error) {
	res, err := c.RedisDB.Get("counter").Int()
	return res, err
}

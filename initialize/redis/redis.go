package redis

import (
	"encoding/json"
	"fmt"
	"time"

	C "github.com/davveo/go-pay/initialize/conf"
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

var Conn *redis.Pool

type R struct {
}

func Init() error {
	Conn = &redis.Pool{
		MaxIdle:     C.GetSettings().Redis.MaxIdle,
		MaxActive:   C.GetSettings().Redis.MaxActive,
		IdleTimeout: C.GetSettings().Redis.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf(
				"%s:%d", C.GetSettings().Redis.Host,
				C.GetSettings().Redis.Port))
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	logrus.Info("Redis初始化成功")

	return nil
}

func (r *R) Set(key string, data interface{}, time int) error {
	conn := Conn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}

	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}

	return nil
}

func (r *R) Exists(key string) bool {
	conn := Conn.Get()
	defer conn.Close()

	exists, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}

	return exists
}

func (r *R) Get(key string) ([]byte, error) {
	conn := Conn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func (r *R) Delete(key string) (bool, error) {
	conn := Conn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func (r *R) LikeDeletes(key string) error {
	conn := Conn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = r.Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}

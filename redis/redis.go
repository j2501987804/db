package redis

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

var (
	pool *redis.Pool
)

// Start 启动redis链接
func Start(url string) {
	pool = newPool(url)

	// ping
	conn := pool.Get()
	if conn != nil {
		log.Printf("redis =>[%s] established.", url)
	}
}

// Get 获取redis链接
func Get() redis.Conn {
	return pool.Get()
}

func newPool(addr string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   100,
		Wait:        true,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			log.Println("open redis...")
			return redis.Dial("tcp", addr)
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}

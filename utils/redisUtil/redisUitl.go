package redisUtil

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

const (
	RedisURL            = "redis://10.0.0.134:6379"
	redisMaxIdle        = 3   //最大空闲连接数
	redisIdleTimeoutSec = 240 //最大空闲连接时间
	RedisPassword       = "2wer43@WER$#"
)

//
//var (
//	conn redis.Conn
//	err error
//)
//
//func init()  {
//	conn, err = redis.Dial("tcp", "106.12.93.100:6379")
//	if err != nil {
//		fmt.Println("Connect to redis err, ", err)
//		return
//	}
//	defer conn.Close()
//	conn.Do("AUTH", "Passw0rd")
//}

//func SetToken(key, value, exp string)  {
//	fmt.Println(key, value, exp)
//	_, err := conn.Do("SET", key, value, "EX", exp)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}

func NewRedisPool(redisURL string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: redisIdleTimeoutSec * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			if err != nil {
				return nil, fmt.Errorf("redis connection error: %s", err)
			}
			//验证redis密码
			if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
				return nil, fmt.Errorf("redis auth password error: %s", authErr)
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			if err != nil {
				return fmt.Errorf("ping redis error: %s", err)
			}
			return nil
		},
	}
}

func set(k, v string) {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("SET", k, v)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func SetKeyWithExp(k, v, exp string) {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("SET", k, v, "EX", exp)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func GetStringValue(k string) string {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	username, err := redis.String(c.Do("GET", k))
	if err != nil {
		fmt.Println("Get Error: ", err.Error())
		return ""
	}
	return username
}

func SetKeyExpire(k string, ex int) {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("EXPIRE", k, ex)
	if err != nil {
		fmt.Println("set error", err.Error())
	}
}

func CheckKey(k string) bool {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	exist, err := redis.Bool(c.Do("EXISTS", k))
	if err != nil {
		fmt.Println(err)
		return false
	} else {
		return exist
	}
}

func DelKey(k string) error {
	c := NewRedisPool(RedisURL).Get()
	defer c.Close()
	_, err := c.Do("DEL", k)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

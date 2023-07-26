package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

type RedisConfig struct {
	Ip       string
	Port     string
	DataBase uint8
	Username string
	Password string
}

type RedisConn struct {
}

func (r RedisConn) connect(config *RedisConfig) {

	dial, err := redis.Dial("tcp", config.Ip+":"+config.Port,
		redis.DialPassword(config.Password),
		redis.DialDatabase(int(config.DataBase)))
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	pool = &redis.Pool{
		// Maximum number of connections allocated by the pool at a given time.
		// When zero, there is no limit on the number of connections in the pool.
		//最大活跃连接数，0代表无限
		MaxActive: 888,
		//最大闲置连接数
		// Maximum number of idle connections in the pool.
		MaxIdle: 20,
		//闲置连接的超时时间
		// Close connections after remaining idle for this duration. If the value
		// is zero, then idle connections are not closed. Applications should set
		// the timeout to a value less than the server's timeout.
		IdleTimeout: time.Second * 100,
		//定义拨号获得连接的函数
		// Dial is an application supplied function for creating and configuring a
		// connection.
		//
		// The connection returned from Dial must not be in a special state
		// (subscribed to pubsub channel, transaction started, ...).
		Dial: func() (redis.Conn, error) {
			return dial, err
		},
	}

}

func (r RedisConn) Close() {
	if pool != nil {
		err := pool.Close()
		if err != nil {
			return
		}
	}
}

func (r RedisConn) GetConn() redis.Conn {
	if pool != nil {
		return pool.Get()
	}
	return nil
}

func NewRedis() RedisConn {
	r := RedisConn{}
	redisConfig := &RedisConfig{
		Ip:       "127.0.0.1",
		Port:     "6379",
		DataBase: 0,
		Username: "",
		Password: "",
	}
	r.connect(redisConfig)
	return r
}

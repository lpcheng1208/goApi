package initialize

import (
	"github.com/go-redis/redis"
	"goApi/global"
	"time"
)

func Redis() {
	redisCfg := global.SERVER_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:         redisCfg.Addr,
		Password:     redisCfg.Password, // no password set
		DB:           redisCfg.DB,       // use default DB
		PoolSize:     redisCfg.RedisPoolSize,
		MaxRetries:   2,
		DialTimeout:  5 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		MinIdleConns: 50,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		global.LOGGER.Sugar().Error(err)
		panic(err)
	} else {
		global.LOGGER.Sugar().Debugf("redis connect ping response:%s", pong)
		global.LOGGER.Debug("init redis success")
		global.REDIS_CONN = client
	}
}

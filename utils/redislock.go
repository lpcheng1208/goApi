package utils

import (
	"errors"
	"github.com/go-redis/redis"
	uuid "github.com/satori/go.uuid"
	"goApi/global"
	"time"
)

/*
基于redis特性的 分布式锁，redis客户端驱动采用go-redis 实现
 */

func GetLock(lockName string, acquireTimeout, lockTimeOut time.Duration) (string, error) {
	code := uuid.NewV4().String()
	fwRedisClient := global.REDIS_CONN

	// 获取毫秒级的时间戳
	endTime := time.Now().Add(acquireTimeout).UnixNano()/1e6
	nowTime := time.Now().UnixNano()/1e6

	for nowTime <= endTime {
		if success, err := fwRedisClient.SetNX(lockName, code, lockTimeOut).Result(); err != nil && err != redis.Nil {
			return "", err
		} else if success {
			return code, nil
		} else if fwRedisClient.TTL(lockName).Val() == -1 { //-2:失效；-1：无过期；
			fwRedisClient.Expire(lockName, lockTimeOut)
		}
		time.Sleep(time.Millisecond)
	}
	return "", errors.New("timeout")
}


//var count = 0  // test assist
func ReleaseLock(lockName, code string) bool {
	fwRedisClient := global.REDIS_CONN
	txf := func(tx *redis.Tx) error {
		if v, err := tx.Get(lockName).Result(); err != nil && err != redis.Nil {
			return err
		} else if v == code {
			_, err := tx.Pipelined(func(pipe redis.Pipeliner) error {
				pipe.Del(lockName)
				return nil
			})
			return err
		}
		return nil
	}

	for {
		if err := fwRedisClient.Watch(txf, lockName); err == nil {
			return true
		} else if err == redis.TxFailedErr {
			global.LOGGER.Sugar().Info("watch key is modified, retry to release lock. err:", err.Error())
		} else {
			global.LOGGER.Sugar().Info("err:", err.Error())
			return false
		}
	}
}
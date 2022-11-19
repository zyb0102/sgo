package cache

import (
	"fmt"
	"testing"
)

func TestNewRedis(t *testing.T) {
	opt := RedisOpt{
		Host:     "47.105.53.133",
		Port:     "6379",
		Password: "a123321666",
		DB:       0,
	}
	redisCli, err := NewRedis(&opt)
	fmt.Println(redisCli, err)
}

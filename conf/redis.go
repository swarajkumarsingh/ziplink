package conf

import (
	"os"
	"time"
)

/*
Redis Configurations
*/

const FreedomRedisTTL = time.Hour * 24

func GetRedisAddr() string {
	return os.Getenv("REDIS_HOST")
}

func getRedisAddrTemp() string {
	return "host.docker.internal:6379"
}

var RedisConf = map[string]interface{}{
	"Addr": getRedisAddrTemp(),
	"SSL":  ENV == ENV_PROD,
	"Username": os.Getenv("REDIS_USER"),
	"Password": os.Getenv("REDIS_PASSWORD"),
}

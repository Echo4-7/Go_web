package redis

import (
	"Web_app/settings"
	"fmt"
	"github.com/go-redis/redis"
)

// 声明一个全局的Rdb变量
var client *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			//viper.GetString("redis.host"),
			//viper.GetInt("redis.port"),
			cfg.Host,
			cfg.Port,
		),
		//Password: viper.GetString("redis.password"),
		//DB:       viper.GetInt("redis.db"),
		//PoolSize: viper.GetInt("redis.poof_size"),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})
	_, err = client.Ping().Result()
	return err
}

func Close() {
	_ = client.Close()
}

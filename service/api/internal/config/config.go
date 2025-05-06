package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RedisCache redis.RedisConf `json:"RedisCache" yaml:"RedisCache"`

	WxMiniApp struct {
		AppId     string
		AppSecret string
	}
	Postgres struct {
		DataSource string
	}
}

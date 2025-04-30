package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	RedisCache cache.CacheConf

	WxMiniApp struct {
		AppId     string
		AppSecret string
	}
	Postgres struct {
		DataSource string
	}
}

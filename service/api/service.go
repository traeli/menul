package main

import (
	"flag"
	"fmt"

	"menul-service/service/api/internal/config"
	"menul-service/service/api/internal/handler"
	"menul-service/service/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"menul-service/service/cache"
)

var configFile = flag.String("f", "./etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	cache.InitRedis(c.RedisCache.Host, c.RedisCache.Pass)
	defer cache.CloseRedis()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package svc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"menul-service/service/api/internal/config"
	"menul-service/service/dao"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  *dao.Query // UserModel 是数据库模型
	FoodModel  *dao.Query // FoodModel
	OrderModel *dao.Query
	DBEngin    *gorm.DB      // GORM 数据库连接
	Redis      *redis.Client //redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	fmt.Println("config:", c.Name)
	fmt.Println("c.Postgres.DataSource", c.Postgres.DataSource)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  c.Postgres.DataSource,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true,
		Logger:                                   logger.Default.LogMode(logger.Info),
	})
	redisClient := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
	if err != nil {
		panic(err)
	}

	// 连接成功后，初始化 ServiceContext
	return &ServiceContext{
		Config:    c,
		UserModel: dao.Use(db), // 初始化 UserModel
		DBEngin:   db,
		FoodModel: dao.Use(db),
		Redis:     redisClient,
	}
}

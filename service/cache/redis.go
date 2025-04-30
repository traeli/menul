package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

const (
	UserTokenKey = "user_token:%s"
)

var Rdb *redis.Client

func InitRedis(Host, Pass string) {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     Host,
		Password: Pass,
		DB:       1,
	})
}

func CloseRedis() {
	err := Rdb.Close()
	if err != nil {
		return
	}
}

func GetToken(userID string) (string, error) {
	token, err := Rdb.Get(context.Background(), fmt.Sprintf(UserTokenKey, userID)).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

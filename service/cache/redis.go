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
		Addr: Host,
		DB:   0,
	})
	fmt.Println("redis init success....", Host)

}

func CloseRedis() {
	err := Rdb.Close()
	if err != nil {
		return
	}
}

func GetToken(token string) (string, error) {
	key := fmt.Sprintf(UserTokenKey, token)
	fmt.Println("key------------------", key)
	userID, err := Rdb.Get(context.Background(), key).Result()
	if err != nil {
		return "", err
	}
	return userID, nil
}

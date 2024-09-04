package dao

import (
	"context"
	"douya/components"
	"douya/consts"
	"douya/model"
	"log"
)

func QueryHello() (string, error) {
	var name string
	db := components.GetMysqlClient(context.Background())
	err := db.Model(&model.User{}).Debug().Select("name").Scan(&name).Error
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "hello " + name, nil
}

func GetNameCache(ctx context.Context) (string, error) {
	data, err := components.GetRedisClient(ctx).Get(ctx, consts.UserNameCacheKey).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func SetNameCache(ctx context.Context, name string) error {
	return components.GetRedisClient(ctx).Set(ctx, consts.UserNameCacheKey, name, consts.UserNameCacheKeyExpireTime).Err()
}

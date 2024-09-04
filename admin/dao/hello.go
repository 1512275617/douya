package dao

import (
	"context"
	"douya/consts"
	"douya/model"
	"douya/pkg"
	"log"
)

func QueryHello() (string, error) {
	var name string
	db := pkg.GetMysqlClient(context.Background())
	err := db.Model(&model.User{}).Debug().Select("name").Scan(&name).Error
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return "hello " + name, nil
}

func GetNameCache(ctx context.Context) (string, error) {
	data, err := pkg.GetRedisClient(ctx).Get(ctx, consts.UserNameCacheKey).Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func SetNameCache(ctx context.Context, name string) error {
	return pkg.GetRedisClient(ctx).Set(ctx, consts.UserNameCacheKey, name, consts.UserNameCacheKeyExpireTime).Err()
}

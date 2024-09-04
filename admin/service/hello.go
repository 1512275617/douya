package service

import (
	"context"
	"douya/dao"
	"errors"
	"github.com/go-redis/redis/v8"
	"log"
)

func Hello(ctx context.Context) (string, error) {
	data, err := dao.GetNameCache(ctx)
	// redis服务异常
	if err != nil && !errors.Is(err, redis.Nil) {
		log.Fatal(err)
		return "", err
	}
	// hit cache return
	if err == nil {
		return data, nil
	}
	//
	name, err := dao.QueryHello()
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	// write cache
	err = dao.SetNameCache(ctx, name)
	if err != nil {
		log.Fatal(err)
	}
	return name, nil
}

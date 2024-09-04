package components

import (
	"github.com/patrickmn/go-cache"
	"time"
)

var localCache *cache.Cache

func LocalCacheInit() {
	localCache = cache.New(5*time.Minute, 10*time.Minute)
}

func GetCache() *cache.Cache {
	return localCache
}

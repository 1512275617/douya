package components

func Init() func() {
	// 初始化mysql
	MysqlInit()

	// 初始化redis
	f1 := RedisInit()

	// 初始化localCache
	LocalCacheInit()

	return func() {
		f1()
	}
}

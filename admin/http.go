package admin

import (
	"context"
	"douya/components"
	"douya/config"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func GetServer() {
	// 初始化gin
	g := gin.New()

	// 请求日志
	g.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.Layout),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	g.Use(gin.Recovery())

	// 加载配置
	config.Init()

	// 初始化路由
	InitRouter(g)

	// 初始化数据库
	components.Init()

	// 启动http服务
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.DefaultConfig.GinCfg.AdminPort),
		Handler: g,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("douya admin start server at :%s ...\n", config.DefaultConfig.GinCfg.AdminPort)

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

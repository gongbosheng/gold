package main

import (
	"context"
	"fmt"
	config "gold/config"
	common "gold/public/common"
	route "gold/route"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// 初始化配置文件
	config.InitConfig()

	// 初始化日志
	common.InitLogger()

	// 初始化数据库
	common.InitDB()

	// 初始化路由
	r := route.InitRouter()

	host := "0.0.0.0"
	port := config.Conf.System.Port

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			common.Log.Fatalf("listen: %s\n", err)
		}
	}()

	common.Log.Info(fmt.Sprintf("Server is running at %s:%d/%s", host, port, config.Conf.System.UrlPathPrefix))

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	<-quit
	common.Log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		common.Log.Fatal("Server forced to shutdown:", err)
	}

	common.Log.Info("Server exiting!")

}

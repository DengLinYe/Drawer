package main

import (
	"Drawer/models"
	"Drawer/routes"
	"Drawer/utils"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	utils.InitLogger()
	utils.Logger.Info("应用程序初始化", "步骤", "日志")

	models.InitDB()
	utils.Logger.Info("应用程序初始化", "步骤", "数据库")

	r := routes.SetupRouter()
	utils.Logger.Info("应用程序初始化", "步骤", "路由")

	srv := &http.Server{
		Addr:    ":8081",
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			utils.Logger.Error("服务器启动失败", "error", err)
			os.Exit(1)
		}

	}()

	utils.Logger.Info("服务器已启动", "地址", srv.Addr)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	utils.Logger.Info("服务器正在关闭...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		utils.Logger.Error("服务器关闭失败", "error", err)
	}
	utils.Logger.Info("服务器已关闭")
}

package main

import (
    "github.com/wh131462/blog-server/internal/handler"
    "github.com/wh131462/blog-server/pkg/config"
    "github.com/gin-gonic/gin"
)

func main() {
    // 加载配置
    cfg := config.Load("config.yaml")

    // 初始化Gin
    r := gin.Default()

    // 注册路由
    handler.RegisterRoutes(r)

    // 启动服务
    r.Run(":" + cfg.Server.Port)
}
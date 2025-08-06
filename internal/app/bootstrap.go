package app

import (
	"fmt"
	v1 "server/internal/api/v1"
	config "server/internal/app/config"

	"github.com/gin-gonic/gin"
)

func Start() {
	// 加载配置文件
	err := config.LoadConfig()
	if err != nil {
		panic("加载配置文件失败")
	}

	// 创建Gin引擎
	r := gin.Default()

	// 初始化所有模块（包括JWT）
	err = InitializeAll(r)
	if err != nil {
		panic("初始化模块失败: " + err.Error())
	}

	// 初始化路由
	v1.SetupRoutes(r, Engine)

	// 启动服务
	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		panic("启动服务失败: " + err.Error())
	}
}

package app

import (
	"fmt"
	config "server/configs"
	v1 "server/internal/api/v1"

	"github.com/gin-gonic/gin"
)

func Start() {
	// 加载配置文件
	err := config.LoadConfig()
	if err != nil {
		panic("加载配置文件失败")
	}

	// 初始化所有模块
	err = InitializeAll()
	if err != nil {
		panic("初始化模块失败")
	}

	r := gin.Default()
	v1.SetupRoutes(r, Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		panic("启动服务失败")
	}

}

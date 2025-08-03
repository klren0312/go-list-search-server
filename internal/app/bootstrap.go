package app

import (
	"fmt"
	v1 "server/internal/api/v1"
	config "server/internal/app/config"

	jwt "github.com/appleboy/gin-jwt/v2"
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
	// jwt 中间件
	authMiddleware, err := jwt.New(initJwtParams())
	if err != nil {
		panic("初始化jwt中间件失败")
	}
	// 初始化中间件
	r.Use(handlerMiddleWare(authMiddleware))

	// 初始化路由
	v1.SetupRoutes(r, Engine)

	err = r.Run(fmt.Sprintf(":%d", config.Conf.App.Port))
	if err != nil {
		panic("启动服务失败")
	}

}

func handlerMiddleWare(authMiddleware *jwt.GinJWTMiddleware) gin.HandlerFunc {
	return func(context *gin.Context) {
		errInit := authMiddleware.MiddlewareInit()
		if errInit != nil {
			panic("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
		}
	}
}

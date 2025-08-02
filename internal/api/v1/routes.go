package v1

import (
	"server/internal/controller"
	"server/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"
)

func SetupRoutes(r *gin.Engine, engine *xorm.Engine) {
	// 定义用户路由
	user := r.Group("/user")
	{
		// 创建 UserService 实例
		UsersService := service.NewUsersService(engine)
		// 创建 UserController 实例
		UsersController := controller.NewUsersController(UsersService)

		user.GET("/", UsersController.GetUsers)
	}
}

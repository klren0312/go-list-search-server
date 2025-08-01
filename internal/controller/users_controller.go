package controller

import (
	"net/http"
	"server/internal/service"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	UsersService *service.UsersService
}

func NewUsersController(UsersService *service.UsersService) *UsersController {
	return &UsersController{
		UsersService: UsersService,
	}
}

func (uc *UsersController) GetUsers(c *gin.Context) {
	users, err := uc.UsersService.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "获取用户列表失败",
			"data": nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "获取用户列表成功",
		"data": gin.H{
			"users": users,
		},
	})
}

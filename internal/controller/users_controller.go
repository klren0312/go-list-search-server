package controller

import (
	"net/http"
	"server/internal/model"
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

func (uc *UsersController) UserLogin(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "参数绑定失败",
			"data": nil,
		})
		return
	}
	
	// 验证用户名和密码
	user, isValid, err := uc.UsersService.VerifyPassword(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "登录失败，服务器错误",
			"data": nil,
		})
		return
	}
	
	// 用户不存在或密码错误
	if user == nil || !isValid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code": 1,
			"msg":  "用户名或密码错误",
			"data": nil,
		})
		return
	}
	
	// 登录成功
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "登录成功",
		"data": gin.H{
			"user": map[string]interface{}{
				"id":         user.Id,
				"user_id":    user.UserId,
				"username":   user.Username,
				"type":       user.Type,
				"reseller_id": user.ResellerId,
			},
		},
	})
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

// CreateUser 创建用户
func (uc *UsersController) CreateUser(c *gin.Context) {
	var req struct {
		UserId     string `json:"user_id"`
		Username   string `json:"username"`
		Password   string `json:"password"`
		Type       int    `json:"type"`
		ResellerId string `json:"reseller_id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "参数绑定失败",
			"data": nil,
		})
		return
	}

	user := &model.Users{
		UserId:     req.UserId,
		Username:   req.Username,
		Password:   req.Password,
		Type:       req.Type,
		ResellerId: req.ResellerId,
	}

	err := uc.UsersService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 1,
			"msg":  "创建用户失败",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建用户成功",
		"data": gin.H{
			"user": user,
		},
	})
}

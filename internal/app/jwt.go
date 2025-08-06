package app

import (
	"server/internal/service"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var identityKey = "id"
var userService *service.UsersService

type User struct {
	Username string
}

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// GetUserService 获取用户服务实例
func GetUserService() *service.UsersService {
	if userService == nil {
		userService = service.NewUsersService(Engine)
	}
	return userService
}

// InitJWT 初始化JWT中间件
func InitJWT(r *gin.Engine) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(initJwtParams())
	if err != nil {
		return nil, err
	}

	// 注册JWT相关路由
	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/refresh_token", authMiddleware.RefreshHandler)

	// 使用JWT中间件保护的路由组
	auth := r.Group("/")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/hello", func(c *gin.Context) {
			claims := jwt.ExtractClaims(c)
			user, _ := c.Get(identityKey)
			c.JSON(200, gin.H{
				"username": claims[identityKey],
				"user":     user.(*User),
				"message":  "Hello, welcome to use JWT!",
			})
		})
	}

	return authMiddleware, nil
}

func initJwtParams() *jwt.GinJWTMiddleware {
	return &jwt.GinJWTMiddleware{
		Realm:       "order system",       // 领域
		Key:         []byte("secret key"), // 密钥
		Timeout:     time.Hour * 24,       // 令牌有效期24小时
		MaxRefresh:  time.Hour * 24 * 7,   // 令牌可刷新7天
		IdentityKey: identityKey,
		PayloadFunc: payloadFunc(),

		IdentityHandler: identityHandler(),
		Authenticator:   authenticator(),
		Authorizator:    authorizator(),
		Unauthorized:    unauthorized(),
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}
}

/**
 * @description: 自定义payload
 * @return {*}
 */
func payloadFunc() func(data interface{}) jwt.MapClaims {
	return func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*User); ok {
			return jwt.MapClaims{
				identityKey: v.Username,
			}
		}
		return jwt.MapClaims{}
	}
}

/**
 * @description: 自定义identityHandler
 * @return {*}
 */
func identityHandler() func(c *gin.Context) interface{} {
	return func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &User{
			Username: claims[identityKey].(string),
		}
	}
}

/**
 * @description: 自定义authenticator
 * @return {*}
 */
func authenticator() func(c *gin.Context) (interface{}, error) {
	return func(c *gin.Context) (interface{}, error) {
		var loginVals login
		if err := c.ShouldBind(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}
		userID := loginVals.Username
		password := loginVals.Password

		// 使用用户服务验证用户名和密码
		userService := GetUserService()
		user, isValid, err := userService.VerifyPassword(userID, password)
		if err != nil {
			return nil, jwt.ErrFailedAuthentication
		}

		if isValid && user != nil {
			return &User{
				Username: userID,
			}, nil
		}

		return nil, jwt.ErrFailedAuthentication
	}
}

func authorizator() func(data interface{}, c *gin.Context) bool {
	return func(data interface{}, c *gin.Context) bool {
		if v, ok := data.(*User); ok {
			// 根据用户名获取用户信息
			userService := GetUserService()
			user, err := userService.GetUserByUsername(v.Username)
			if err != nil || user == nil {
				return false
			}

			return user.Type >= 0
		}
		return false
	}
}

func unauthorized() func(c *gin.Context, code int, message string) {
	return func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	}
}

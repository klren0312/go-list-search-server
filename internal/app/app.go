package app

import (
	"fmt"
	
	"github.com/gin-gonic/gin"
)

// InitializeAll 初始化所有模块
func InitializeAll(r *gin.Engine) error {
	// 初始化MySQL
	err := InitializeMySQL()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}

	// 初始化JWT
	_, err = InitJWT(r)
	if err != nil {
		return fmt.Errorf("JWT初始化错误: %v", err)
	}

	return nil
}

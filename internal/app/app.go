package app

import (
	"fmt"
)

// InitializeAll 初始化所有模块
func InitializeAll() error {
	err := InitializeMySQL()
	if err != nil {
		return fmt.Errorf("MySQL初始化错误: %v", err)
	}

	return nil
}

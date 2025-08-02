package repository

import (
	"server/internal/model"

	"github.com/go-xorm/xorm"
)

type UsersRepository struct {
	engine *xorm.Engine
}

func NewUsersRepository(engine *xorm.Engine) *UsersRepository {
	return &UsersRepository{engine: engine}
}

// GetUsers 获取所有用户
func (r *UsersRepository) GetUsers() ([]*model.Users, error) {
	var users []*model.Users
	err := r.engine.Table((&model.Users{}).TableName()).Find(&users)
	return users, err
}

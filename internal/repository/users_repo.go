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

// GetUserByUsername 根据用户名获取用户
func (r *UsersRepository) GetUserByUsername(username string) (*model.Users, error) {
	user := &model.Users{}
	has, err := r.engine.Table((&model.Users{}).TableName()).Where("username = ?", username).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return user, nil
}

// CreateUser 创建用户
func (r *UsersRepository) CreateUser(user *model.Users) error {
	_, err := r.engine.Table((&model.Users{}).TableName()).Insert(user)
	return err
}

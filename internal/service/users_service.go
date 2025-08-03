package service

import (
	"server/internal/model"
	"server/internal/repository"
	"server/internal/utils"

	"github.com/go-xorm/xorm"
)

type UsersService struct {
	userRepo *repository.UsersRepository
}

func NewUsersService(engine *xorm.Engine) *UsersService {
	return &UsersService{
		userRepo: repository.NewUsersRepository(engine),
	}
}

func (us *UsersService) GetUsers() ([]*model.Users, error) {
	return us.userRepo.GetUsers()
}

// GetUserByUsername 根据用户名获取用户
func (us *UsersService) GetUserByUsername(username string) (*model.Users, error) {
	return us.userRepo.GetUserByUsername(username)
}

// VerifyPassword 验证用户密码
func (us *UsersService) VerifyPassword(username, password string) (*model.Users, bool, error) {
	// 根据用户名获取用户
	user, err := us.GetUserByUsername(username)
	if err != nil {
		return nil, false, err
	}
	
	// 如果用户不存在
	if user == nil {
		return nil, false, nil
	}
	
	// 验证密码
	isValid := utils.VerifyPassword(password, user.Salt, user.Password)
	
	return user, isValid, nil
}

// CreateUser 创建用户
func (us *UsersService) CreateUser(user *model.Users) error {
	// 生成随机盐值
	salt, err := utils.GenerateSalt(16)
	if err != nil {
		return err
	}
	
	// 保存原始密码
	originalPassword := user.Password
	
	// 设置盐值
	user.Salt = salt
	
	// 对密码进行加盐哈希
	user.Password = utils.HashPassword(originalPassword, salt)
	
	// 保存用户
	return us.userRepo.CreateUser(user)
}

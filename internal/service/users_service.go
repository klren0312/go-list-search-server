package service

import (
	"server/internal/model"
	"server/internal/repository"

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

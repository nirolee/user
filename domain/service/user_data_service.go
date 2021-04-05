package service

import (
	"github.com/nirolee/user.git/domain/model"
	"github.com/nirolee/user.git/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (int64, error)
	DeleteUser(int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, password string) (bool, error)
}

func NewsUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFormPassword([]byte(userPassword), bcrayt.DefaultCost)
}

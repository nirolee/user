package service

import (
	"errors"
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
	return &UserDataService{UserRepository: userRepository}
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

//加密用户密码
func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

//验证用户密码
func ValidatePassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码错误")
	}
	return true, nil
}

//添加用户
func (u *UserDataService)AddUser(user *model.User) (userId int64, err error){
	pwdByte, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return user.Id, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)

}

func (u *UserDataService) DeleteUser (userId int64) (err error) {
	return u.UserRepository.DeleteUserById(userId)

}

//更新用户
func (u *UserDataService)UpdateUser(user *model.User, isChangePwd bool) (err error){
	if isChangePwd {
		pwdByte, err := GeneratePassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

//根据用户名称查找用户信息
func (u *UserDataService)FindUserByName(userName string) (user *model.User, err error) {
	return u.FindUserByName(userName)
}

func (u *UserDataService)CheckPwd(userName string, password string) (isOk bool, err error){
	user, err := u.UserRepository.FindUserByName(userName)
	if err!= nil{
		return false, err
	}
	return ValidatePassword(password, user.HashPassword)
}


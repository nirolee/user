package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/nirolee/user.git/domain/model"
)

type IUserRepository interface {
	//初始化数据表
	InitTable() error
	//根据用户名称搜索用户信息
	FindUserByName(string) (*model.User, error)
	//根据用户id搜索用户信息
	FindUserById(int64) (*model.User, error)
	//创建用户
	CreateUser(*model.User) (int64, error)
	//根据用户id删除用户
	DeleteUserById(int64) error
	//更新用户信息
	UpdateUser(*model.User) error
	//查找所有用户
	FindAll() ([]model.User, error)
}

//创建UserRepository
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{mysqlDb: db}
}

type UserRepository struct {
	mysqlDb *gorm.DB
}

//初始化表
func (u *UserRepository) InitTable() error {
	return u.mysqlDb.CreateTable(&model.User{}).Error
}

func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

func (u *UserRepository) FindUserById(userId int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDb.First(user, userId).Error
}

func (u *UserRepository) CreateUser(user *model.User) (userId int64, err error) {
	user = &model.User{}
	return user.Id, u.mysqlDb.Create(user).Error
}

func (u *UserRepository) DeleteUserById(userId int64) error {
	return u.mysqlDb.Where("id = ?", userId).Delete(&model.User{}).Error
}

func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}

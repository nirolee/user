package repository

import "github.com/nirolee/user.git/domain/model"

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

func NewUserRepository(db *gorm.Db)  {
	
}

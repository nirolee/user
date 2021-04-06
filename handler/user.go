package handler

import (
	"context"
	"github.com/nirolee/user.git/domain/model"
	"github.com/nirolee/user.git/domain/service"
	user "github.com/nirolee/user.git/proto/user"
)

type User struct{
	service.UserDataService
}

//注册
func (u *User)Register(ctx context.Context, in *user.UserRegisterRequest,
	out *user.UserRegisterResponse) error  {
	userRegister := &model.User{
		UserName: in.UserName,
		FirstName: in.UserName,
		HashPassword: in.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegister)
	if err != nil {
		return err
	}
	out.Message = "添加成功"
	return nil
}

//登录
func (u *User)Login(ctx context.Context, in *user.UserLoginRequest,
	out *user.UserLoginResponse) error  {
	isOk, err := u.UserDataService.CheckPwd(in.UserName, in.Pwd)
	if err != nil {
		return err
	}
	out.IsSuccess = isOk
	return nil
}

//查询用户信息
func (u *User)GetUserInfo(ctx context.Context, in *user.UserInfoRequest,
	out *user.UserInfoResponse) error  {
	userInfor, err := u.UserDataService.FindUserByName(in.UserName)
	if err != nil {
		return err
	}
	out = UserForResponse(userInfor)
	return nil
}

func UserForResponse(userModel *model.User) (response *user.UserInfoResponse) {
	response.UserName = userModel.UserName
	response.FirstName = userModel.FirstName
	response.UserId = userModel.Id
	return response

}
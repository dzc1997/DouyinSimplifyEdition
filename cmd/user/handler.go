package main

import (
	"context"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/user/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/user/service"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
)

type UserServiceImpl struct {}

func (s *UserServiceImpl) UserLogin(ctx context.Context, req *user.UserLoginRequest) (resp *user.UserLoginResponse, err error) {
	resp = new(user.UserLoginResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserLoginResp(errno.ParamErr)
		return resp, nil
	}

	uid, err := service.NewLoginUserService(ctx).CheckUser(req)
	if err != nil {
		resp = pack.BuildUserLoginResp(err)
		return resp, nil
	}
	resp.UserId = uid
	resp = pack.BuildUserLoginResp(errno.Success)
	return resp, nil
}

func (s *UserServiceImpl) UserRegister(ctx context.Context, req *user.UserRegisterRequest) (resp *user.UserRegisterResponse, err error) {
	resp = new(user.UserRegisterResponse)

	if len(req.Username) == 0 || len(req.Password) == 0 {
		resp = pack.BuildUserRegisterResp(errno.ParamErr)
	}

	userId, err := service.NewRegisterUserService(ctx).RegisterUser(req)
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return
	}

	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: userId,
	})
	if err != nil {
		resp = pack.BuildUserRegisterResp(err)
		return
	}

	resp = pack.BuildUserRegisterResp(errno.Success)
	resp.UserId = userId
	resp.Token = token
	return resp, nil
}

func (s *UserServiceImpl) UserInfo(ctx context.Context, req *user.UserRequest) (resp *user.UserResponse, err error) {
	resp = new(user.UserResponse)

	if req.UserId == 0 {
		resp = pack.BuildUserInfoResp(errno.ParamErr)
	}

	user_, err := service.NewUserInfoService(ctx).UserInfo(req)
	if err != nil {
		resp = pack.BuildUserInfoResp(err)
		return
	}

	resp = pack.BuildUserInfoResp(errno.Success)
	resp.User = user_
	return resp, nil
}
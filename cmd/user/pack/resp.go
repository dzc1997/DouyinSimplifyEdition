package pack

import (
	"context"
	"errors"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)

func userLoginResp(err errno.ErrNo) *user.UserLoginResponse {
	return &user.UserLoginResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildUserLoginResp(err error) *user.UserLoginResponse {
	if err == nil {
		return userLoginResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userLoginResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return userLoginResp(s)
}

func userRegisterResp(err errno.ErrNo) *user.UserRegisterResponse {
	return &user.UserRegisterResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildUserRegisterResp(err error) *user.UserRegisterResponse {
	if err == nil {
		return userRegisterResp(errno.Success)
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userRegisterResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return userRegisterResp(s)
}

func userInfoResp(err errno.ErrNo) *user.UserResponse {
	return &user.UserResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildUserInfoResp(err error) *user.UserResponse {
	if err == nil {
		return userInfoResp(errno.Success)
	}

	klog.CtxInfof(context.TODO(), "BuildUserInfoResp err %v", err)

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return userInfoResp(e)
	}
	s := errno.ServiceErr.WithMessage(err.Error())
	return userInfoResp(s)
}

package service

import (
	"context"
	"crypto/md5"
	"fmt"
	"io"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/user/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)

type LoginUserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginUserService {
	return &LoginUserService{
		ctx: ctx,
	}
}

func (s *LoginUserService) CheckUser(req *user.UserLoginRequest) (int64, error) {
	h := md5.New()
	if _, err := io.WriteString(h, req.Password); err != nil {
		return 0, err
	}
	passWord := fmt.Sprintf("%x", h.Sum(nil))

	userName := req.Username
	klog.Infof( "CheckUser got req %+v. passwd md5: %v\n", req, passWord)
	users, err := db.QueryUserByName(s.ctx, userName)
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errno.UserNotExistErr
	}
	u := users[0]

	klog.CtxInfof(context.TODO(), "CheckUser got user %+v.", u)

	if u.Password != passWord {
		return 0, errno.LoginErr
	}
	return int64(u.ID), nil
}

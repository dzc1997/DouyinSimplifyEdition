package service

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
)

func TestUserInfo(t *testing.T) {
	type reqArgs struct {
		userId int64
		token  string
	}
	tests := []struct {
		name    string
		args    reqArgs
		wantErr bool
	}{
		{
			name: "用户信息",
			args: reqArgs{
				userId: 1,
				token:  Token,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user_, err := NewUserInfoService(context.Background()).UserInfo(&user.UserRequest{
				UserId: tt.args.userId,
				Token:  tt.args.token,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("UserInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(user_)
			klog.Info(tt.name + " success")
		})
	}
}

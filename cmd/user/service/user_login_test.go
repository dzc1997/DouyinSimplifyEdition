package service

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
)

func TestCheckUser(t *testing.T) {
	type reqArgs struct {
		username string
		password string
	}
	tests := []struct {
		name    string
		args    reqArgs
		wantErr bool
	}{
		{
			name: "用户登陆",
			args: reqArgs{
				username: "dance",
				password: "qq789456",
			},
			wantErr: false,
		},
		{
			name: "测试登陆不存在的用户",
			args: reqArgs{
				username: "UnExist",
				password: "UnExist",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewLoginUserService(context.Background()).CheckUser(&user.UserLoginRequest{
				Username: tt.args.username,
				Password: tt.args.password,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(tt.name + " success")
		})
	}
}

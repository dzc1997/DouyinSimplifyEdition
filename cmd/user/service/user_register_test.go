package service

import (
	"bytes"
	"context"
	"io"
	"os"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/user/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/user"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
)

var Token string

func TestMain(m *testing.M) {

	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	token, err := Jwt.CreateToken(jwt.CustomClaims{
		Id: int64(1),
	})
	if err != nil {
		klog.Errorf("create token fail, %v", err.Error())
		panic(err)
	}
	Token = token

	db.Init()
	oss.Init()

	path := oss.Path + "/public/test.mp4"
	file, err := os.Open(path)
	if err != nil {
		klog.Errorf("open local file %v fail", path)
		panic(err)
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		panic(err)
	}

	m.Run()
}

func TestRegisterUser(t *testing.T) {
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
			name: "用户注册",
			args: reqArgs{
				username: "QQQ",
				password: "hhh123",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userId, err := NewRegisterUserService(context.Background()).RegisterUser(&user.UserRegisterRequest{
				Username: tt.args.username,
				Password: tt.args.password,
			})
			if (err != nil) != tt.wantErr {
				t.Errorf("RegisterUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(tt.name + " success")
			klog.Info(userId)
		})
	}
}

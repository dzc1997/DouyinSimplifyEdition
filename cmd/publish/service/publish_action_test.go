package service

import (
	"bytes"
	"context"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/publish/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
	"io"
	"os"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
)

var File []byte
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
	File = buf.Bytes()
	m.Run()
}

func TestPublishAction(t *testing.T) {
	type reqArgs struct {
		token string
		data  []byte
		title string
	}
	tests := []struct {
		name    string
		args    reqArgs
		wantErr bool
	}{
		{
			name: "视频投稿功能",
			args: reqArgs{
				token: Token,
				data:  File,
				title: "测试",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewPublishService(context.Background()).Publish(&publish.PublishActionRequest{Token: tt.args.token, Title: tt.args.title, Data: tt.args.data})
			if (err != nil) != tt.wantErr {
				t.Errorf("%v fail, %v, wantErr %v", tt.name, err, tt.wantErr)
				return
			}
			klog.Info(tt.name + " success")
		})
	}
}

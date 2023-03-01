package service

import (
	"bytes"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/favorite/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/favorite"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
	"io"
	"os"
	"testing"
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
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		panic(err)
	}
	m.Run()
}

func TestFavoriteAction(t *testing.T) {
	type reqArgs struct {
		token      string
		videoId    int64
		actionType int32
	}
	favoriteAcitonTest := []struct {
		name   string
		args   reqArgs
		result bool
	}{
		{
			name: "点赞操作",
			args: reqArgs{
				token:      Token,
				videoId:    1,
				actionType: 1,
			},
			result: false,
		},
		{
			name: "取消赞操作",
			args: reqArgs{
				token:      Token,
				videoId:    1,
				actionType: 2,
			},
			result: false,
		},
	}

	for _, tt := range favoriteAcitonTest {
		t.Run(tt.name, func(t *testing.T) {
			err := NewFavoriteActionService(context.Background()).FavoriteAction(&favorite.FavoriteActionRequest{
				Token:      tt.args.token,
				VideoId:    tt.args.videoId,
				ActionType: tt.args.actionType,
			})
			if (err != nil) != tt.result {
				t.Errorf("FavoriteAction() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(tt.name + " success")
		})
	}
}

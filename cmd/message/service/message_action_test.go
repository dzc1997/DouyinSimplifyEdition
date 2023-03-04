package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/message/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/message"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
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
	m.Run()
}

func TestMessageAction(t *testing.T) {
	type reqArgs struct {
		token      string
		toUserId   int64
		actionType int32
		content    string
	}
	messageActionTest := []struct {
		name   string
		args   reqArgs
		result bool
	}{
		{
			name: "发送消息",
			args: reqArgs{
				token:      Token,
				toUserId:   2,
				actionType: 1,
				content:    "测试",
			},
			result: false,
		},
	}
	for _, tt := range messageActionTest {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewActionService(context.Background()).MessageAction(&message.MessageActionRequest{
				Token:      tt.args.token,
				ToUserId:   tt.args.toUserId,
				ActionType: tt.args.actionType,
				Content:    tt.args.content,
			})
			if (err != nil) != tt.result {
				t.Errorf("MessageAction() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(tt.name + "success")
		})
	}
}

package service

import (
	"bytes"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/comment/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/comment"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
	"io"
	"os"
	"testing"
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
	}
	defer file.Close()
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		panic(err)
	}
	File = buf.Bytes()
	m.Run()
}

func TestCommentAction(t *testing.T) {
	type reqArgs struct {
		token       string
		videoId     int64
		actionType  int32
		commentText string
		commentId 	int64
	}
	createTest := []struct {
		name   string
		args   reqArgs
		result bool
	}{
		{
			name: "评论操作",
			args: reqArgs{
				token:       Token,
				videoId:     1,
				actionType:  1,
				commentText: "测试",
			},
			result: true,
		},
		{
			name: "删除评论",
			args : reqArgs{
				token: Token,
				videoId: 1,
				actionType: 2,
				commentId: 1,
			},
		},
	}

	for _, tt := range createTest {
		t.Run(tt.name, func(t *testing.T) {
			comment__, err := NewCommentActionService(context.Background()).CommentAction(&comment.CommentActionRequest{
				Token:       tt.args.token,
				VideoId:     tt.args.videoId,
				ActionType:  tt.args.actionType,
				CommentText: &tt.args.commentText,
				CommentId: &tt.args.commentId,
			})
			if (err != nil) != tt.result {
				t.Errorf("CommentAction() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(comment__)
			klog.Info(tt.name + " success")
		})
	}
}

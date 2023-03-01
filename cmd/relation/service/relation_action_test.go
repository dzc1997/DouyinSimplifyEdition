package service

import (
	"bytes"
	"context"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/oss"
	"io"
	"os"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
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

func TestRelationAction(t *testing.T) {
	type reqArgs struct {
		token      string
		toUserId   int64
		actionType int64
	}
	tests := []struct {
		name    string
		args    reqArgs
		wantErr bool
	}{
		{
			name: "关注操作",
			args: reqArgs{
				token:      Token,
				toUserId:   2,
				actionType: 1,
			},
			wantErr: false,
		},
		{
			name: "取消关注操作",
			args: reqArgs{
				token:      Token,
				toUserId:   2,
				actionType: 2,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewRelationActionService(context.Background()).RelationAction(&relation.RelationActionRequest{Token: tt.args.token, ToUserId: tt.args.toUserId, ActionType: int32(tt.args.actionType)})
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationAction() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(tt.name + " success")
		})
	}
}

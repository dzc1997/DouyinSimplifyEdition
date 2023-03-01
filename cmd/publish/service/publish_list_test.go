package service

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish"
)

func TestPublishListAction(t *testing.T) {
	type args struct {
		token  string
		userId int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "视频列表",
			args: args{
				token:  Token,
				userId: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			videoList, err := NewPublishListService(context.Background()).PublishList(&publish.PublishListRequest{UserId: tt.args.userId, Token: tt.args.token})
			if (err != nil) != tt.wantErr {
				t.Errorf("PublishList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(tt.name + "success")
			klog.Info(videoList)
		})
	}
}

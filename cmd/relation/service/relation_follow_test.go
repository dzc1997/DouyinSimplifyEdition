package service

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
)

func TestFollowList(t *testing.T) {
	type reqArgs struct {
		token  string
		userId int64
	}
	tests := []struct {
		name    string
		args    reqArgs
		wantErr bool
	}{
		{
			name: "关注列表",
			args: reqArgs{
				token:  Token,
				userId: 1,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewFollowListService(context.Background()).FollowList(&relation.RelationFollowListRequest{Token: tt.args.token, UserId: tt.args.userId})
			if (err != nil) != tt.wantErr {
				t.Errorf("FollowList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			klog.Info(tt.name + " success")
		})
	}
}

package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/favorite"
	"testing"
)

func TestCommentList(t *testing.T) {
	type reqArgs struct {
		userId int64
		token  string
	}
	tests := []struct {
		name    string
		reqArgs reqArgs
		result  bool
	}{
		{
			name: "正常操作",
			reqArgs: reqArgs{
				userId: 1,
				token:  Token,
			},
			result: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			favoriteList, err := NewFavoriteListService(context.Background()).FavoriteList(&favorite.FavoriteListRequest{
				UserId: tt.reqArgs.userId,
				Token:  tt.reqArgs.token,
			})
			if (err != nil) != tt.result {
				t.Errorf("favoriteList() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(favoriteList)
			klog.Info(tt.name + " success")
		})
	}
}

package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/comment"
	"testing"
)

func TestCommentList(t *testing.T) {
	type reqArgs struct {
		token   string
		videoId int64
	}
	tests := []struct {
		name    string
		reqArgs reqArgs
		result  bool
	}{
		{
			name: "正常操作",
			reqArgs: reqArgs{
				token:   Token,
				videoId: 1,
			},
			result: false,
		},
		{
			name: "测试评论列表的不正确视频id",
			reqArgs: reqArgs{
				token:   Token,
				videoId: 99999999,
			},
			result: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			commentList, err := NewCommentListService(context.Background()).CommentList(&comment.CommentListRequest{
				Token: tt.reqArgs.token,
				VideoId: tt.reqArgs.videoId,
			})
			if (err != nil) != tt.result {
				t.Errorf("CommentList() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(commentList)
			klog.Info(tt.name + " success")
		})
	}
}

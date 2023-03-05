package service

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/message"
	"testing"
)

func TestMessageChat(t *testing.T) {
	type reqArgs struct {
		token    string
		toUserId int64
	}
	messageChatTest := []struct {
		name   string
		args   reqArgs
		result bool
	}{
		{
			name: "聊天记录",
			args: reqArgs{
				token:    Token,
				toUserId: 2,
			},
			result: false,
		},
	}
	for _, tt := range messageChatTest {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMessageChatService(context.Background()).MessageChat(&message.MessageChatRequest{
				Token:    tt.args.token,
				ToUserId: tt.args.toUserId,
			})
			if (err != nil) != tt.result {
				t.Errorf("MessageChat() error = %v, result %v", err, tt.result)
				return
			}
			klog.Info(tt.name + "success")
		})
	}
}

package service

import (
	"context"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/message/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/message/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/message"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
)

type MessageChatService struct {
	ctx context.Context
}

func NewMessageChatService(ctx context.Context) *MessageChatService {
	return &MessageChatService{ctx: ctx}
}

func (s *MessageChatService) MessageChat(req *message.MessageChatRequest) ([]*message.Message, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)
	messages := make([]*db.MessageRaw, 0)
	messages, err := db.QueryMessageById(s.ctx, currentId, req.ToUserId)
	if err != nil {
		return nil, err
	}
	messages_ := make([]*message.Message, 0)
	messages_ = pack.MessageList(messages)

	return messages_, nil
}

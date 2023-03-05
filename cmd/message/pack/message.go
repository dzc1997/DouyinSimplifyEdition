package pack

import (
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/message/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/message"
)

func Chat(messageRaw *db.MessageRaw) message.Message {
	message_ := message.Message{
		Id:         int64(messageRaw.ID),
		ToUserId:   messageRaw.ToUserID,
		FromUserId: messageRaw.FromUserID,
		Content:    messageRaw.Content,
	}
	return message_
}

func MessageList(messages []*db.MessageRaw) []*message.Message {
	messageList := make([]*message.Message, 0)
	var timeLayoutStr = "2006-01-02 15:04:05"
	for _, message_ := range messages {
		messageList = append(messageList, &message.Message{
			Id:         int64(message_.ID),
			ToUserId:   message_.ToUserID,
			FromUserId: message_.FromUserID,
			Content:    message_.Content,
			CreateTime: message_.CreateTime.Format(timeLayoutStr),
		})
	}
	return messageList
}

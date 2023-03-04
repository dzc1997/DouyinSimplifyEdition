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

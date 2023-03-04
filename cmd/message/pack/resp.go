package pack

import (
	"errors"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/message"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)

func BuildMessageActionResp(err error) *message.MessageActionResponse {
	if err == nil {
		return MessageActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return MessageActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return MessageActionResp(s)
}

func MessageActionResp(err errno.ErrNo) *message.MessageActionResponse {
	return &message.MessageActionResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildMessageChatRep(err error) *message.MessageChatResponse {
	if err == nil {
		return messageChatResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return messageChatResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return messageChatResp(s)
}

func messageChatResp(err errno.ErrNo) *message.MessageChatResponse {
	return &message.MessageChatResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

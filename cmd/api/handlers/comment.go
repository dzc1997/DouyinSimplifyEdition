package handlers

import (
	"context"
	"strconv"
	"unicode/utf8"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/api/rpc"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/comment"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")

	klog.Infof("CommentAction token[%v] videoID[%v] action[%v]\n", token, videoIdStr, actionTypeStr)

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	if actionType == constants.AddComment {
		commentText := c.Query("comment_text")

		if length := utf8.RuneCountInString(commentText); length > 20 {
			SendResponse(c, errno.CommentTextErr)
			return
		}

		req := &comment.CommentActionRequest{Token: token, VideoId: videoId, CommentText: &commentText, ActionType: constants.AddComment}
		comment_, err := rpc.CommentAction(context.Background(), req)
		if err != nil {
			SendResponse(c, errno.ConvertErr(err))
			return
		}
		SendCommentActionResponse(c, errno.Success, comment_)

	} else if actionType == constants.DelComment {
		commentIdStr := c.Query("comment_id")

		commentId, err := strconv.ParseInt(commentIdStr, 10, 64)
		if err != nil {
			SendResponse(c, errno.ParamParseErr)
		}

		req := &comment.CommentActionRequest{Token: token, VideoId: videoId, CommentId: &commentId, ActionType: constants.DelComment}
		comment_, err := rpc.CommentAction(context.Background(), req)
		if err != nil {
			SendResponse(c, errno.ConvertErr(err))
			return
		}
		SendCommentActionResponse(c, errno.Success, comment_)

	} else {
		SendResponse(c, errno.ParamErr)
	}
}

func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")

	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
	}

	req := &comment.CommentListRequest{Token: token, VideoId: videoId}
	commentList, err := rpc.CommentList(context.Background(), req)
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	SendCommentListResponse(c, errno.Success, commentList)
}

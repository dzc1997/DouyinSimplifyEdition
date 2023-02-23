package handlers

import (
	"context"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/api/rpc"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/favorite"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionTypeStr := c.Query("action_type")

	klog.Infof("FavoriteAction token[%v] videoID[%v] action[%v]\n", token, videoIdStr, actionTypeStr)

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

	err = rpc.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
		Token: token, VideoId: videoId, ActionType: int32(actionType),
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, errno.Success)
}

//FavoriteList get favorite list info
func FavoriteList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	token := c.Query("token")

	klog.Infof("FavoriteList param %+v userID %v token %v\n", c.Request.URL, userIdStr, token)

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	videoList, err := rpc.FavoriteList(context.Background(), &favorite.FavoriteListRequest{
		Token: token, UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	SendFavoriteListResponse(c, errno.Success, videoList)
}

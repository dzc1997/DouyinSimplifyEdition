package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"strconv"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/api/rpc"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")

	klog.Infof("PublishAction title[%v] token[%v]\n", title, token)

	data, _, err := c.Request.FormFile("data")
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}
	defer data.Close()

	if length := len(title); length <= 0 || length > 128 {
		SendResponse(c, errno.ParamErr)
		return
	}

	//处理视频数据
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, data); err != nil {
		SendResponse(c, errno.VideoDataCopyErr)
		return
	}
	video := buf.Bytes()

	err = rpc.PublishVideoData(context.Background(), &publish.PublishActionRequest{
		Token: token,
		Title: title, Data: video,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	SendResponse(c, errno.Success)
}

//PublishList get publish list
func PublishList(c *gin.Context) {
	token := c.Query("token")
	userIdStr := c.Query("user_id")

	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		SendResponse(c, errno.ParamParseErr)
		return
	}

	klog.Infof("QueryVideoList req userID[%v] token[%v]", userIdStr, token)

	videoList, err := rpc.QueryVideoList(context.Background(), &publish.PublishListRequest{
		Token:  token,
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err))
		return
	}

	klog.Infof("QueryVideoList got videoList len[%v]", len(videoList))

	if len(videoList) > 0 {
		marshal, _ := json.Marshal(videoList)
		klog.Infof("QueryVideoList got videoList [%s]", string(marshal))
	}

	SendPublishListResponse(c, errno.Success, videoList)
}

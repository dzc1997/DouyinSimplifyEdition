package handlers

import (
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type UserResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserId     int64  `json:"user_id,omitempty"`
	Token      string `json:"token,omitempty"`
}

type UserInfoResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	User       interface{} `json:"user"`
}

func SendResponse(c *gin.Context, err error) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, Response{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
	})
}

func SendUserResponse(c *gin.Context, err error, userId int64, token string) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		UserId:     userId,
		Token:      token,
	})
}

func SendUserInfoResponse(c *gin.Context, err error, user interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, UserInfoResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		User:       user,
	})
}

type FeedResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  interface{} `json:"video_list,omitempty"`
	NextTime   int64       `json:"next_time,omitempty"`
}

func SendFeedResponse(c *gin.Context, err error, videoList interface{}, nextTime int64) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, FeedResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videoList,
		NextTime:   nextTime,
	})
}

type PublishListResponse struct {
	StatusCode int32       `json:"status_code"`
	StatusMsg  string      `json:"status_msg"`
	VideoList  interface{} `json:"video_list,omitempty"`
}


func SendPublishListResponse(c *gin.Context, err error, videoList interface{}) {
	Err := errno.ConvertErr(err)
	c.JSON(http.StatusOK, PublishListResponse{
		StatusCode: Err.ErrCode,
		StatusMsg:  Err.ErrMsg,
		VideoList:  videoList,
	})
}

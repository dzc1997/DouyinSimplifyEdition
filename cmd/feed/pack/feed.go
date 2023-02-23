package pack

import (
	"time"

	"github.com/dzc1997/DouyinSimplifyEdition/cmd/feed/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/feed"
)

func VideoInfo(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, favoriteMap map[int64]*db.FavoriteRaw, relationMap map[int64]*db.RelationRaw) ([]*feed.Video, int64) {
	videoList := make([]*feed.Video, 0)
	var nextTime int64
	for _, video := range videoData {
		videoUser, ok := userMap[video.UserId]
		if !ok {
			videoUser = &db.UserRaw{}
			videoUser.ID = 0
		}

		var isFavorite = false
		var isFollow = false

		if currentId != -1 {
			_, ok := favoriteMap[int64(video.ID)]
			if ok {
				isFavorite = true
			}
			_, ok = relationMap[video.UserId]
			if ok {
				isFollow = true
			}
		}

		if len(videoUser.Name) == 0 {
			videoUser.Name = "未知用户"
		}

		if len(videoUser.Signature) == 0 {
			videoUser.Signature = "这个用户很懒, 什么都没写"
		}

		videoList = append(videoList, &feed.Video{
			Id: int64(video.ID),
			Author: &feed.User{
				Id:              int64(videoUser.ID),
				Name:            videoUser.Name,
				FollowCount:     &videoUser.FollowCount,
				FollowerCount:   &videoUser.FollowerCount,
				IsFollow:        isFollow,
				Avatar:          &videoUser.Avatar,
				BackgroundImage: &videoUser.BackgroundImage,
				Signature:       &videoUser.Signature,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	if len(videoData) == 0 {
		nextTime = time.Now().UnixMilli()
	} else {
		nextTime = videoData[len(videoData)-1].UpdatedAt.UnixMilli()
	}

	return videoList, nextTime
}

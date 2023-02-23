package pack

import (
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/publish/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish"
)

func PublishInfo(currentId int64, videoData []*db.VideoRaw, userMap map[int64]*db.UserRaw, favoriteMap map[int64]*db.FavoriteRaw, relationMap map[int64]*db.RelationRaw) []*publish.Video {
	videoList := make([]*publish.Video, 0)
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

		videoList = append(videoList, &publish.Video{
			Id: int64(video.ID),
			Author: &publish.User{
				Id:              int64(videoUser.ID),
				Name:            videoUser.Name,
				FollowCount:     &videoUser.FollowCount,
				FollowerCount:   &videoUser.FollowerCount,
				IsFollow:        isFollow,
				Avatar:          &videoUser.Avatar,
				BackgroundImage: &videoUser.BackgroundImage,
				Signature:       &videoUser.Signature,
				TotalFavorited:  &videoUser.TotalFavorited,
				WorkCount:       &videoUser.WorkCount,
				FavoriteCount:   &videoUser.FavoriteCount,
			},
			PlayUrl:       video.PlayUrl,
			CoverUrl:      video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount:  video.CommentCount,
			IsFavorite:    isFavorite,
			Title:         video.Title,
		})
	}

	return videoList
}

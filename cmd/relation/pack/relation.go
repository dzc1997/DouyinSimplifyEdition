package pack

import (
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
)

func UserList(currentId int64, users []*db.UserRaw, relationMap map[int64]*db.RelationRaw) []*relation.User {
	userList := make([]*relation.User, 0)
	for _, user := range users {
		var isFollow bool = false

		if currentId != -1 {
			_, ok := relationMap[int64(user.ID)]
			if ok {
				isFollow = true
			}
		}
		userList = append(userList, &relation.User{
			Id:            int64(user.ID),
			Name:          user.Name,
			FollowCount:   &user.FollowCount,
			FollowerCount: &user.FollowerCount,
			IsFollow:      isFollow,
		})
	}
	return userList
}

func FriendList(users []*db.UserRaw) []*relation.FriendUser {
	userList := make([]*relation.FriendUser, 0)
	for _, user := range users {
		userList = append(userList, &relation.FriendUser{
			Id:              int64(user.ID),
			Name:            user.Name,
			FollowCount:     &user.FollowCount,
			FollowerCount:   &user.FollowerCount,
			IsFollow:        true,
			Avatar:          &user.Avatar,
			BackgroundImage: &user.BackgroundImage,
			Signature:       &user.Signature,
			TotalFavorited:  &user.TotalFavorited,
			WorkCount:       &user.WorkCount,
			FavoriteCount:   &user.FavoriteCount,
		})
	}
	return userList
}

package service

import (
	"context"
	"errors"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
	"sort"
)

type FriendListService struct {
	ctx context.Context
}

func NewFriendListService(ctx context.Context) *FriendListService {
	return &FriendListService{ctx: ctx}
}

func (s *FriendListService) FriendList(req *relation.RelationFriendListRequest) ([]*relation.FriendUser, error) {
	var friendIds []int64
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)
	if currentId != req.UserId {
		return nil, errors.New("token check fail")
	}

	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("userId not exist")
	}

<<<<<<< HEAD
	
=======
>>>>>>> 649cdfe (dy)
	follows, err := db.QueryFollowById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	followIds := make([]int64, 0)
	for _, follow := range follows {
		followIds = append(followIds, follow.ToUserId)
	}

	followers, err := db.QueryFollowerById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	followerIds := make([]int64, 0)
	for _, follower := range followers {
		followerIds = append(followerIds, follower.UserId)
	}

<<<<<<< HEAD

=======
>>>>>>> 649cdfe (dy)
	sort.Slice(followIds, func(i, j int) bool {
		return followIds[i] < followIds[j]
	})
	sort.Slice(followerIds, func(i, j int) bool {
		return followerIds[i] < followerIds[j]
	})

<<<<<<< HEAD
	
=======
>>>>>>> 649cdfe (dy)
	n, m := len(followIds), len(followerIds)
	for i, j := 0, 0; i < n && j < m; i++ {
		for followerIds[j] < followIds[i] {
			j++
		}
		if followIds[i] == followerIds[j] {
			friendIds = append(friendIds, followIds[i])
			j++
		}
	}
	friends, err := db.QueryUserByIds(s.ctx, friendIds)
	if err != nil {
		return nil, err
	}
	friendList := pack.FriendList(friends)
	return friendList, nil
}

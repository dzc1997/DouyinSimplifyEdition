package service

import (
	"context"
	"errors"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/dal/db"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/jwt"
)

type FollowListService struct {
	ctx context.Context
}

// NewFollowListService new FollowListService
func NewFollowListService(ctx context.Context) *FollowListService {
	return &FollowListService{ctx: ctx}
}

// FollowList get user follow list info
func (s *FollowListService) FollowList(req *relation.RelationFollowListRequest) ([]*relation.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("userId not exist")
	}

	//获取目标用户关注的用户id号
	relations, err := db.QueryFollowById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	userIds := make([]int64, 0)
	for _, relation_ := range relations {
		userIds = append(userIds, relation_.ToUserId)
	}

	//获取用户id号对应的用户信息
	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}

	var relationMap map[int64]*db.RelationRaw
	if currentId == -1 {
		relationMap = nil
	} else {
		//获取当前用户和这些用户的关注信息
		relationMap, err = db.QueryRelationByIds(s.ctx, currentId, userIds)
		if err != nil {
			return nil, err
		}
	}
	userList := pack.UserList(currentId, users, relationMap)
	return userList, nil
}


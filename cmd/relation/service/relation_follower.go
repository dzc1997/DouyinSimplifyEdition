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

type FollowerListService struct {
	ctx context.Context
}

func NewFollowerListService(ctx context.Context) *FollowerListService {
	return &FollowerListService{ctx: ctx}
}

func (s *FollowerListService) FollowerList(req *relation.RelationFollowerListRequest) ([]*relation.User, error) {
	Jwt := jwt.NewJWT([]byte(constants.SecretKey))
	currentId, _ := Jwt.CheckToken(req.Token)

	user, err := db.QueryUserByIds(s.ctx, []int64{req.UserId})
	if err != nil {
		return nil, err
	}
	if len(user) == 0 {
		return nil, errors.New("userId not exist")
	}

	relations, err := db.QueryFollowerById(s.ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	userIds := make([]int64, 0)
	for _, relation_ := range relations {
		userIds = append(userIds, relation_.UserId)
	}

	users, err := db.QueryUserByIds(s.ctx, userIds)
	if err != nil {
		return nil, err
	}

	var relationMap map[int64]*db.RelationRaw
	if currentId == -1 {
		relationMap = nil
	} else {
		relationMap, err = db.QueryRelationByIds(s.ctx, currentId, userIds)
		if err != nil {
			return nil, err
		}
	}

	userList := pack.UserList(currentId, users, relationMap)
	return userList, nil
}

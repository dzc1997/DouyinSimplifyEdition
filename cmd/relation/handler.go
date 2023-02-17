package main

import (
	"context"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/relation/service"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)


type RelationServiceImpl struct{}

func (s *RelationServiceImpl) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest) (resp *relation.RelationFriendListResponse, err error){
	resp = new(relation.RelationFriendListResponse)
	return resp, nil
}

func (s *RelationServiceImpl) RelationAction(ctx context.Context, req *relation.RelationActionRequest) (resp *relation.RelationActionResponse, err error) {
	resp = new(relation.RelationActionResponse)

	if len(req.Token) == 0 || req.ToUserId == 0 || req.ActionType == 0 {
		resp = pack.BuildRelationActionResp(errno.ParamErr)
		return resp, nil
	}

	err = service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		resp = pack.BuildRelationActionResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationActionResp(errno.Success)
	return resp, nil
}


func (s *RelationServiceImpl) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest) (resp *relation.RelationFollowListResponse, err error) {
	resp = new(relation.RelationFollowListResponse)

	if req.UserId == 0 {
		resp = pack.BuildRelationFollowListResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewFollowListService(ctx).FollowList(req)
	if err != nil {
		resp = pack.BuildRelationFollowListResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationFollowListResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}

func (s *RelationServiceImpl) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) (resp *relation.RelationFollowerListResponse, err error) {
	resp = new(relation.RelationFollowerListResponse)

	if req.UserId == 0 {
		resp = pack.BuildRelationFollowerListResp(errno.ParamErr)
		return resp, nil
	}

	userList, err := service.NewFollowerListService(ctx).FollowerList(req)
	if err != nil {
		resp = pack.BuildRelationFollowerListResp(err)
		return resp, nil
	}
	resp = pack.BuildRelationFollowerListResp(errno.Success)
	resp.UserList = userList
	return resp, nil
}

package pack

import (
	"errors"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)

func BuildRelationActionResp(err error) *relation.RelationActionResponse {
	if err == nil {
		return relationActionResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationActionResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationActionResp(s)
}

func relationActionResp(err errno.ErrNo) *relation.RelationActionResponse {
	return &relation.RelationActionResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildRelationFollowListResp(err error) *relation.RelationFollowListResponse {
	if err == nil {
		return relationFollowListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFollowListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationFollowListResp(s)
}

func relationFollowListResp(err errno.ErrNo) *relation.RelationFollowListResponse {
	return &relation.RelationFollowListResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}

func BuildRelationFollowerListResp(err error) *relation.RelationFollowerListResponse {
	if err == nil {
		return relationFollowerListResp(errno.Success)
	}

	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return relationFollowerListResp(e)
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return relationFollowerListResp(s)
}

func relationFollowerListResp(err errno.ErrNo) *relation.RelationFollowerListResponse {
	return &relation.RelationFollowerListResponse{StatusCode: err.ErrCode, StatusMsg: &err.ErrMsg}
}
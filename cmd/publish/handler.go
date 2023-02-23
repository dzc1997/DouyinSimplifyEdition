package main

import (
	"context"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/publish/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/publish/service"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)

type PublishServiceImpl struct{}

// PublishAction implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishAction(ctx context.Context, req *publish.PublishActionRequest) (resp *publish.PublishActionResponse, err error) {
	resp = new(publish.PublishActionResponse)

	if len(req.Token) == 0 || len(req.Title) == 0 || req.Data == nil {
		resp = pack.BuildPublishActionResp(errno.ParamErr)
		return resp, nil
	}

	klog.Infof("PublishAction token[%v] title[%v] datalen[%v]", req.Token, req.Title, len(req.Data))

	err = service.NewPublishService(ctx).Publish(req)
	if err != nil {
		resp = pack.BuildPublishActionResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishActionResp(errno.Success)
	return resp, nil
}

// PublishList implements the PublishServiceImpl interface.
func (s *PublishServiceImpl) PublishList(ctx context.Context, req *publish.PublishListRequest) (resp *publish.PublishListResponse, err error) {
	resp = new(publish.PublishListResponse)

	if req.UserId <= 0 {
		resp = pack.BuildPublishListResp(errno.ParamErr)
		return resp, nil
	}

	videoList, err := service.NewPublishListService(ctx).PublishList(req)
	if err != nil {
		resp = pack.BuildPublishListResp(err)
		return resp, nil
	}
	resp = pack.BuildPublishListResp(errno.Success)
	resp.VideoList = videoList
	return resp, nil
}

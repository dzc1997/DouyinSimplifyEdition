package main

import (
	"context"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/feed/pack"
	"github.com/dzc1997/DouyinSimplifyEdition/cmd/feed/service"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/feed"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
)



type FeedServiceImpl struct{}

// Feed implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) Feed(ctx context.Context, req *feed.FeedRequest) (resp *feed.FeedResponse, err error) {
	resp = new(feed.FeedResponse)

	if *req.LatestTime <= 0 {
		resp = pack.BuildFeedResp(errno.ParamErr)
		return resp, nil
	}

	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)
	if err != nil {
		resp = pack.BuildFeedResp(err)
		return resp, nil
	}
	resp = pack.BuildFeedResp(errno.Success)
	resp.VideoList = videos
	resp.NextTime = &nextTime
	return resp, nil
}


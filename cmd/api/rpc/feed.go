package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/feed"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/feed/feedservice"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var feedClient feedservice.Client

func initFeedRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := feedservice.NewClient(
		constants.FeedServiceName,
		client.WithMiddleware(middleware.CommonMiddleware),
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	feedClient = c
}

//Feed query list of video info
func Feed(ctx context.Context, req *feed.FeedRequest) ([]*feed.Video, int64, error) {
	resp, err := feedClient.Feed(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode != 0 {
		return nil, 0, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.VideoList, *resp.NextTime, nil
}

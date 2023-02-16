package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/publish/publishservice"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var publishClient publishservice.Client

func initPublishRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := publishservice.NewClient(
		constants.PublishServiceName,
		client.WithInstanceMW(middleware.ClientMiddleware),
		client.WithMuxConnection(1),
		client.WithRPCTimeout(10*time.Second),
		client.WithConnectTimeout(10000*time.Millisecond),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithSuite(trace.NewDefaultClientSuite()),
		client.WithResolver(r),
	)
	if err != nil {
		panic(err)
	}
	publishClient = c
}

// PublishVideoData upload video data
func PublishVideoData(ctx context.Context, req *publish.PublishActionRequest) error {
	resp, err := publishClient.PublishAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return nil
}

// QueryVideoList get a list of video releases
func QueryVideoList(ctx context.Context, req *publish.PublishListRequest) ([]*publish.Video, error) {
	resp, err := publishClient.PublishList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.VideoList, nil
}


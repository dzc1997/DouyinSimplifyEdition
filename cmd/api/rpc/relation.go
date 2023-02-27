package rpc

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
	"github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation/relationservice"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/constants"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/errno"
	"github.com/dzc1997/DouyinSimplifyEdition/pkg/middleware"
	etcd "github.com/kitex-contrib/registry-etcd"
	trace "github.com/kitex-contrib/tracer-opentracing"
	"time"
)

var relationClient relationservice.Client

func initRelationRpc() {
	r, err := etcd.NewEtcdResolver([]string{constants.EtcdAddress})
	if err != nil {
		panic(err)
	}

	c, err := relationservice.NewClient(
		constants.RelationServiceName,
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
	relationClient = c
}

func RelationAction(ctx context.Context, req *relation.RelationActionRequest) error {
	resp, err := relationClient.RelationAction(ctx, req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 0 {
		return errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return nil
}

func FriendList(ctx context.Context, req *relation.RelationFriendListRequest) ([]*relation.FriendUser, error)  {
	resp, err := relationClient.RelationFriendList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.UserList, nil
}

func FollowList(ctx context.Context, req *relation.RelationFollowListRequest) ([]*relation.User, error) {
	resp, err := relationClient.RelationFollowList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.UserList, nil
}

func FollowerList(ctx context.Context, req *relation.RelationFollowerListRequest) ([]*relation.User, error) {
	resp, err := relationClient.RelationFollowerList(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 0 {
		return nil, errno.NewErrNo(resp.StatusCode, *resp.StatusMsg)
	}
	return resp.UserList, nil
}


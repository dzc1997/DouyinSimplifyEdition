// Code generated by Kitex v0.4.4. DO NOT EDIT.

package relationservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	relation "github.com/dzc1997/DouyinSimplifyEdition/kitex_gen/relation"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	RelationAction(ctx context.Context, req *relation.RelationActionRequest, callOptions ...callopt.Option) (r *relation.RelationActionResponse, err error)
	RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest, callOptions ...callopt.Option) (r *relation.RelationFollowListResponse, err error)
	RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.RelationFollowerListResponse, err error)
	RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest, callOptions ...callopt.Option) (r *relation.RelationFriendListResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kRelationServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kRelationServiceClient struct {
	*kClient
}

func (p *kRelationServiceClient) RelationAction(ctx context.Context, req *relation.RelationActionRequest, callOptions ...callopt.Option) (r *relation.RelationActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationAction(ctx, req)
}

func (p *kRelationServiceClient) RelationFollowList(ctx context.Context, req *relation.RelationFollowListRequest, callOptions ...callopt.Option) (r *relation.RelationFollowListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowList(ctx, req)
}

func (p *kRelationServiceClient) RelationFollowerList(ctx context.Context, req *relation.RelationFollowerListRequest, callOptions ...callopt.Option) (r *relation.RelationFollowerListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFollowerList(ctx, req)
}

func (p *kRelationServiceClient) RelationFriendList(ctx context.Context, req *relation.RelationFriendListRequest, callOptions ...callopt.Option) (r *relation.RelationFriendListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RelationFriendList(ctx, req)
}

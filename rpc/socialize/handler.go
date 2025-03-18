package main

import (
	socialize "TikTok-rpc/rpc/socialize/kitex_gen/socialize"
	"context"
)

// SocializeServiceImpl implements the last service interface defined in the IDL.
type SocializeServiceImpl struct{}

// Follow implements the SocializeServiceImpl interface.
func (s *SocializeServiceImpl) Follow(ctx context.Context, req *socialize.FollowRequest) (resp *socialize.FollowResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFollowList implements the SocializeServiceImpl interface.
func (s *SocializeServiceImpl) QueryFollowList(ctx context.Context, req *socialize.QueryFollowListRequest) (resp *socialize.QueryFollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFollowerList implements the SocializeServiceImpl interface.
func (s *SocializeServiceImpl) QueryFollowerList(ctx context.Context, req *socialize.QueryFollowerListRequest) (resp *socialize.QueryFollowerListResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryFriendList implements the SocializeServiceImpl interface.
func (s *SocializeServiceImpl) QueryFriendList(ctx context.Context, req *socialize.QueryFriendListRequest) (resp *socialize.QueryFriendListResponse, err error) {
	// TODO: Your code here...
	return
}

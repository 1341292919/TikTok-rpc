package main

import (
	video "TikTok-rpc/rpc/video/kitex_gen/video"
	"context"
)

// VideoServiceImpl implements the last service interface defined in the IDL.
type VideoServiceImpl struct{}

// PublishVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) PublishVideo(ctx context.Context, req *video.PublishRequest) (resp *video.PublishResponse, err error) {
	// TODO: Your code here...
	return
}

// QueryList implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) QueryList(ctx context.Context, req *video.QueryPublishListRequest) (resp *video.QueryPublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// SearchVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) SearchVideo(ctx context.Context, req *video.SearchVideoByKeywordRequest) (resp *video.SearchVideoByKeywordResponse, err error) {
	// TODO: Your code here...
	return
}

// GetPopularVideo implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetPopularVideo(ctx context.Context, req *video.GetPopularListRequest) (resp *video.GetPopularListResponse, err error) {
	// TODO: Your code here...
	return
}

// GetVideoStream implements the VideoServiceImpl interface.
func (s *VideoServiceImpl) GetVideoStream(ctx context.Context, req *video.VideoStreamRequest) (resp *video.VideoStreamResponse, err error) {
	// TODO: Your code here...
	return
}

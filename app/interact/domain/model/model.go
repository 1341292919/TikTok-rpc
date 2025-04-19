package model

type Comment struct {
	Id         int64
	Uid        int64
	Type       int64
	ParentId   int64
	Content    string
	CreateAT   int64
	UpdateAT   int64
	LikeCount  int64
	ChildCount int64
}

// 与video服务中的video模型有些出入
type Video struct {
	Id           string
	Uid          string
	Title        string
	Description  string
	VideoUrl     string
	CoverUrl     string
	CreateAT     string
	UpdateAT     string
	DeleteAT     string
	VisitCount   int64
	CommentCount int64
	LikeCount    int64
}
type InteractReq struct {
	Uid        int64
	VideoId    int64
	CommentId  int64
	PageNum    int64
	PageSize   int64
	ActionType int64
	Content    string
	Type       int64 //0视频 1评论
}
type LikeKey struct {
	Uid       int64
	VideoId   int64
	CommentId int64
	Type      int64 //0 视频 1 评论
	Status    int64 //0点赞 1 不点赞
	Time      int64
}
type VideoLikeCountKey struct {
	Id    int64
	Count int64
}
type UserLike struct {
	Uid       int64
	VideoId   int64
	CommentId int64
	Type      int64 //0 视频 1 评论
}
type CommentLikeCountKey struct {
	Id    int64
	Count int64
}

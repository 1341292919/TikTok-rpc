package model

type Video struct {
	Id           int64
	Uid          int64
	Title        string
	Description  string
	VideoUrl     string
	CoverUrl     string
	CreateAT     int64
	UpdateAT     int64
	DeleteAT     int64
	VisitCount   int64
	CommentCount int64
	LikeCount    int64
}
type VideoReq struct {
	Uid      int64
	Keyword  string
	PageNum  int64
	PageSize int64
	FromDate int64
	ToDate   int64
	Username string
}
type LikeCount struct {
	Id    int64
	Count int64
}

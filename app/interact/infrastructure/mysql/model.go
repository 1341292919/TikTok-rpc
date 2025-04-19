package mysql

import "time"

type UserLike struct {
	UserId   int64
	TargetId int64
	LikedAt  time.Time
	Type     int64
}
type Comment struct {
	Id         int64
	UserId     int64
	Content    string
	ParentId   int64
	Type       int64
	ChildCount int64
	LikeCount  int64
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

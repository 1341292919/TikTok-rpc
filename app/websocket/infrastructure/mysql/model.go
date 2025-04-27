package mysql

import "time"

type Message struct {
	Id        int64
	UserId    int64
	TargetId  int64
	Content   string
	CreatedAt time.Time
	Status    int64
	Type      int64
}

package model

type Message struct {
	Id        int64
	UserId    int64
	TargetId  int64
	Content   string
	CreatedAT int64
	Type      int64
	Status    int64 //0 未读 1 已读
}

type ChatReq struct {
	UserId   int64
	TargetId int64
	PageSize int64
	PageNum  int64
}

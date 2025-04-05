package model

//传给usecase操作的User，实际上也是usecase与db交互的User，但是实际上存在空间的很大浪费

type User struct {
	Uid       int64
	UserName  string
	Password  string
	AvatarUrl string
	OptSecret string
	CreateAT  int64
	UpdateAT  int64
	DeleteAT  int64
	Code      string
}
type MFA struct {
	Secret string
	Qrcode string
}
type MFAMessage struct {
	Status int64
	Secret string
}

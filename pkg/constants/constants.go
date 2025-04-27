package constants

const (
	QiNiuBucket    = "tiktok1341292919"
	QiNiuAccessKey = "BrnJRlH-n-PTi_4M_zT_AvXYFIGQt9xVq-bbYOGh"
	QiNiuSecretKey = "K7j2CR_pRexVKnwJclMqcTavKP3hDM9T2TPGAcrP"
	QiNiuDomain    = "https://portal.qiniu.com/"
)

// redis
const (
	RedisUserName       = "default"
	RedisPassWord       = "Yang"
	RedisHost           = "127.0.0.1"
	RedisPort           = "6379"
	VideoIdKey          = "VideoId"
	VideoKey            = "Video"
	MessageKey          = "Message"
	VideoLikeCountKey   = "VideoLikeCount"
	CommentLikeCountKey = "CommentLikeCount"
)

// Service Name
const (
	UserServiceName      = "user"
	UserETCD             = "127.0.0.1:2379"
	VideoServiceName     = "video"
	VideoETCD            = "127.0.0.1:2379"
	InteractServiceName  = "interact"
	InteractETCD         = "127.0.0.1:2379"
	SocializeETCD        = "127.0.0.1:2379"
	WebsocketServiceName = "websocket"
	WebsocketETCD        = "127.0.0.1:2379"
)

// jwt
const (
	AccessTokenKey  = "AccessToken_key"
	RefreshTokenKey = "RefreshToken_key"
)

// websocket
const (
	PrivateChat    = "to_user_message"
	GroupChat      = "to_group_message"
	PrivateMessage = "get_private_history"
	GroupMessage   = "get_group_history"
)
const (
	ContextUserId   = "user_id"
	AvatarStorePath = "/home/yang/Desktop/resource/avatar"
	VideoStorePath  = "/home/yang/Desktop/resource/video"
	CoverStorePath  = "/home/yang/Desktop/resource/cover"
	MySQLDSN        = "root:casaos@tcp(127.0.0.1:3306)/casaos?charset=utf8mb4&parseTime=true"
	TableUser       = "user"
	TableComment    = "comment"
	TableMessage    = "chat_message"
	TableFollower   = "user_follows"
	TableUserLike   = "user_likes"
	TableVideo      = "video"
)

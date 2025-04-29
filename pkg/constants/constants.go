package constants

// redis
const (
	VideoIdKey = "VideoId"
	VideoKey   = "Video"
	MessageKey = "Message"
)

// Service Name
const (
	UserServiceName      = "user"
	VideoServiceName     = "video"
	InteractServiceName  = "interact"
	WebsocketServiceName = "websocket"
	GatewayServiceName   = "gateway"
	SocializeServiceName = "socialize "
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
	TableUser       = "user"
	TableComment    = "comment"
	TableMessage    = "chat_message"
	TableUserLike   = "user_likes"
	TableVideo      = "video"
)
const (
	RedisDBVideo     = 0
	RedisDBInteract  = 1
	RedisDBWebsocket = 2
)

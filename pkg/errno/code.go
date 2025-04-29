package errno

const (
	SuccessCode = 10000
	SuccessMsg  = "success"
)

const (
	ParamVerifyErrorCode  = 20000 + iota
	ParamMissingErrorCode //参数缺失
	ParamLogicalErrorCode
)

const (
	AuthInvalidCode        = 30000 + iota // 鉴权失败
	AuthAccessExpiredCode                 // 访问令牌过期
	AuthRefreshExpiredCode                // 刷新令牌过期
	AuthNoTokenCode                       // 没有 token
)

const (
	RepeatOperationErrorCode = 40000 + iota
	ServiceUserExistCode
	ServiceUserNotExistCode
	ServiceUserPasswordError
	ServiceVideoNotExist
	ServiceCommentNotExist
	ServiceRepeatOperation
	ServiceNoAuthority
)

const (
	InternalServiceErrorCode  = 50000 + iota // 内部服务错误
	InterFileProcessErrorCode                //文件处理错误
	InternalDatabaseErrorCode
	InternalRedisErrorCode // Redis错误
	InternalRPCErrorCode   //
	InternalWebSocketError
)

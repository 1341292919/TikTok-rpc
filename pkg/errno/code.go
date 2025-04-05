package errno

const (
	SuccessCode = 10000
	SuccessMsg  = "success"
)

// 200xx：参数错误
const (
	ParamVerifyErrorCode  = 20000 + iota
	ParamMissingErrorCode //参数缺失
	ParamLogicalErrorCode
)

// 300xx:鉴权问题
const (
	AuthInvalidCode             = 30000 + iota // 鉴权失败
	AuthAccessExpiredCode                      // 访问令牌过期
	AuthRefreshExpiredCode                     // 刷新令牌过期
	AuthNoTokenCode                            // 没有 token
	AuthNoOperatePermissionCode                // 没有操作权限
	AuthMissingTokenCode                       // 缺少 token
	IllegalOperatorCode                        // 不合格的操作(比如传入 payment status时传入了一个不存在的 status)
)

// 400xx：业务错误
const (
	RepeatOperationErrorCode = 40000 + iota
	ServiceUserExist
	ServiceUserNotExist
	ServiceUserPasswordError
	ServiceVideoNotExist
	ServiceCommentNotExist
)

// 500xx: 内部错误，Internal 打头
// 服务级别的错误, 发生的时候说明我们程序自身出了问题
const (
	InternalServiceErrorCode  = 50000 + iota // 内部服务错误
	InternalDatabaseErrorCode                // 数据库错误
	InterFileProcessErrorCode                //文件处理错误
	InternalRedisErrorCode                   // Redis错误
	InternalNetworkErrorCode                 // 网络错误
	InternalESErrorCode                      // ES错误
	InternalKafkaErrorCode                   // kafka 错误
	OSOperateErrorCode
	IOOperateErrorCode
	InsufficientStockErrorCode
	InternalRPCErrorCode
	InternalRocketmqErrorCode
)

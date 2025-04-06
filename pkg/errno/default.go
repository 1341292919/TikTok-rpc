package errno

var (
	Success              = NewErrNo(SuccessCode, SuccessMsg)
	InternalServiceError = NewErrNo(InternalServiceErrorCode, "internal server error")

	ParamVerifyError  = NewErrNo(ParamVerifyErrorCode, "parameter validation failed")
	ParamMissingError = NewErrNo(ParamMissingErrorCode, "missing parameter")

	AuthInvalid        = NewErrNo(AuthInvalidCode, "authentication failure")
	AuthAccessExpired  = NewErrNo(AuthAccessExpiredCode, "token expiration")
	AuthRefreshExpired = NewErrNo(AuthRefreshExpiredCode, "token refresh expired")
	AuthNoToken        = NewErrNo(AuthNoTokenCode, "lack of token")

	UserIdMissingError = NewErrNo(InternalServiceErrorCode, "get Userid from context failed")
)

package service

import "TikTok-rpc/pkg/errno"

// 用于检验user服务里各种需要的参数
type UserVerifyFuncs func() error

func (svc *UserService) Verify(funcs ...UserVerifyFuncs) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// 这里检验MFAcode是否合规
func (svc *UserService) VerifyMFACode(code string) UserVerifyFuncs {
	return func() error {
		if code == "" {
			return errno.NewErrNo(errno.ParamMissingErrorCode, "required code to verify")
		}
		if len(code) != 6 {
			return errno.NewErrNo(errno.ParamVerifyErrorCode, "code length error")
		}
		return nil
	}
}

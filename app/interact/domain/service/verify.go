package service

import "TikTok-rpc/pkg/errno"

type InteractVerifyFunc func() error

func (svc *InteractService) Verify(funcs ...InteractVerifyFunc) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// 检验action_type
func (svc *InteractService) VerifyActionType(act int64) InteractVerifyFunc {
	return func() error {
		if act == 0 || act == 1 {
			return nil
		}
		return errno.NewErrNo(errno.ParamLogicalErrorCode, "Invalid Action Type")
	}
}
func (svc *InteractService) VerifyPageParam(pagesize, pagenum int64) InteractVerifyFunc {
	return func() error {
		if pagenum <= 0 || pagesize <= 0 {
			return errno.NewErrNo(errno.ParamLogicalErrorCode, "pagesize,pagenum must over 0")
		}
		return nil
	}
}

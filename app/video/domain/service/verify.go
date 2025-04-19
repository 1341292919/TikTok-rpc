package service

import "TikTok-rpc/pkg/errno"

type VideoVerifyFunc func() error

func (svc *VideoService) Verify(funcs ...VideoVerifyFunc) error {
	for _, f := range funcs {
		if err := f(); err != nil {
			return err
		}
	}
	return nil
}

// 检验传入的日期
func (svc *VideoService) VerifyDate(todate, fromdate int64) VideoVerifyFunc {
	return func() error {
		//传入的日期为空
		if todate == 0 && fromdate == 0 {
			return nil
		} else if todate < fromdate {
			return errno.NewErrNo(errno.ParamLogicalErrorCode, "ToDate earlier than FromDate")
		}
		return nil
	}
}
func (svc *VideoService) VerifyPageParam(pagesize, pagenum int64) VideoVerifyFunc {
	return func() error {
		if pagenum <= 0 || pagesize <= 0 {
			return errno.NewErrNo(errno.ParamLogicalErrorCode, "pagesize,pagenum must over 0")
		}
		return nil
	}
}

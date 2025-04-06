package service

import (
	"TikTok-rpc/pkg/constants"
	"TikTok-rpc/pkg/errno"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserIDFromContext(c *app.RequestContext) int64 {
	data := c.Keys[constants.ContextUserId]
	uid, err := convertToInt64(data)
	if err != nil {
		panic(errno.UserIdMissingError)
	}
	return uid
}
func convertToInt64(value interface{}) (int64, error) {
	switch v := value.(type) {
	case int:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("无法转换为int64，类型为 %T", value)
	}
}

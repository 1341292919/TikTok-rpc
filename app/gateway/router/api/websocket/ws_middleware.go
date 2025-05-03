package websock

import "github.com/cloudwego/hertz/pkg/app"

func _homeMW() []app.HandlerFunc {
	return _wsAuth()
}

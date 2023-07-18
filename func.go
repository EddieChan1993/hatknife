package hatknife

import (
	"fmt"
	"reflect"
	"runtime"
)

// GetModName 获取结构体名
func GetModName(i interface{}) string {
	return reflect.TypeOf(i).String()
}

// PanicStack 捕获recover的panic堆栈
func PanicStack() {
	buf := make([]byte, 1<<10)
	runtime.Stack(buf, true)
	fmt.Println(string(buf))
}

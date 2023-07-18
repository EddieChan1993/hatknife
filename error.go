package hatknife

import (
	"fmt"
	"runtime"
)

// PanicStack 捕获recover的panic堆栈
func PanicStack() {
	buf := make([]byte, 1<<10)
	runtime.Stack(buf, true)
	fmt.Println(string(buf))
}

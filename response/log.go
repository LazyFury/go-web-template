package response

import (
	"fmt"
	"log"
	"runtime"
)

// Error Error
func Error(err interface{}) {
	_, file, line, ok := runtime.Caller(1)
	if ok {
		LogError(fmt.Sprintf("path:%s:%d", file, line), "发生错误的文件")
	}
	panic(err)
}

// LogError LogError
func LogError(err interface{}, tips string) {
	log.Printf("\x1b[31m[%s]: %v \x1b[0m\n", tips, err)
}

var (
	ErrCodeTextMap = []ErrorCodeTextInterface{
		ErrorCodeText,
	}
)

func PushErrCodeTextMap(_map ErrorCodeTextInterface) {
	if _map != nil {
		ErrCodeTextMap = append(ErrCodeTextMap, _map)
	}
}

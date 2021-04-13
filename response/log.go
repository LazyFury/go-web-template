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
		LogError(fmt.Sprintf("%s %d", file, line))
	}
	panic(err)
}

// LogError LogError
func LogError(err interface{}) {
	log.Printf("\n\x1b[31m[Custom Debug Result]: %v \x1b[0m\n\n", err)
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

package Common

import (
	"runtime"
	"strings"
)

const _os_windows  = "windows"

func IsWindowsRuntime() bool {
	return strings.ToLower(runtime.GOOS) == _os_windows
}

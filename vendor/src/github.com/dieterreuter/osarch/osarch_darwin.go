// +build darwin
package osarch

import (
	"fmt"
	"runtime"
)

func Arch() string {
	s := fmt.Sprint(runtime.GOARCH)
	return s
}

func OsArch() string {
	s := fmt.Sprint(runtime.GOOS, "/", runtime.GOARCH)
	return s
}

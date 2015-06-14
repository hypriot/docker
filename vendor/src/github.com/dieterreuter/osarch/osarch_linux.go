// +build linux
package osarch

import (
	"fmt"
	"golang.org/x/sys/unix"
	"runtime"
)

func getUnameMachine() string {
	// same as 'uname -m'
	var n unix.Utsname
	_ = unix.Uname(&n)

	s := make([]byte, len(n.Machine))
	var lens int
	for ; lens < len(n.Machine); lens++ {
		if n.Machine[lens] == 0 {
			break
		}
		s[lens] = uint8(n.Machine[lens])
	}
	return string(s[0:lens])
}

func Arch() string {
	s := fmt.Sprint(runtime.GOARCH)
	if runtime.GOARCH == "arm" {
		s = fmt.Sprint(s, "/", getUnameMachine())
	}
	return s
}

func OsArch() string {
	s := fmt.Sprint(runtime.GOOS, "/", runtime.GOARCH)
	if runtime.GOARCH == "arm" {
		s = fmt.Sprint(s, "/", getUnameMachine())
	}
	return s
}

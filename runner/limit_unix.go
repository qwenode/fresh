//go:build !windows
// +build !windows

package runner

import (
	"fmt"
	"syscall"
)

func initLimit() {
	var rLimit syscall.Rlimit
	rLimit.Max = 1000000
	rLimit.Cur = 1000000
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	if err != nil {
		fmt.Println("Error Setting Rlimit ", err)
	}
}

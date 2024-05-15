//go:build linux

package ctime

import (
	"os"
	"syscall"
	"time"
)

func CreateTime(info os.FileInfo) time.Time {
	st := info.Sys().(*syscall.Stat_t)
	return time.Unix(int64(st.Ctim.Sec), int64(st.Ctim.Nsec)) //nolint:unconvert
}

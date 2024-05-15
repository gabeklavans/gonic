//go:build !linux

package ctime

import (
	"os"
	"time"
)

func CreateTime(info os.FileInfo) time.Time {
	return info.ModTime()
}

package utils

import (
	"AdminPro/internal/glog"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var modDir string

func init() {
	stdout, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		return
	}
	fp := string(bytes.TrimSpace(stdout))
	if fp == os.DevNull || fp == "" {
		return
	}
	modDir = filepath.Dir(fp)
}

func GoSafe(handle func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				var buf [4096]byte
				n := runtime.Stack(buf[:], false)
				tmpStr := fmt.Sprintf("err: %v, panic==> %s\n", r, string(buf[:n]))
				glog.Errorf("Recovered from panic in goroutine: %v", tmpStr)
			}
		}()
		handle()
	}()
}

// Path 兼容返回相对执行目录
func Path(fp string) string {
	if modDir == "" || fp == "" || filepath.IsAbs(fp) {
		return fp
	}
	return filepath.Join(modDir, fp)
}

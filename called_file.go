package helpers

import (
	"runtime"
	"errors"
	"strings"
	"path/filepath"
)

func GetCalledFileNameSkip(skip int, abs... bool) string {
	_, filename, _, ok := runtime.Caller(skip)
	if !ok {
		panic(errors.New("Information unavailable."))
	}
	if len(abs) == 0 || !abs[0] {
		filename = strings.TrimPrefix(filename, filepath.Join(GOPATH, "src"))
		return filename[1:]
	}
	return filename
}

func GetCalledDir(abs ...bool) string {
	file := GetCalledFileNameSkip(2, abs...)
	return filepath.Dir(file)
}
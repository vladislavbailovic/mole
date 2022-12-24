package internal

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetTestFilePath(relpath string) string {
	_, b, _, _ := runtime.Caller(0)
	cwd := filepath.Dir(b)
	pth := filepath.Join(cwd, "../testdata", relpath)
	return pth
}

func GetTestFile(relpath string) []byte {
	pth := GetTestFilePath(relpath)
	buffer, err := os.ReadFile(pth)
	if err != nil {
		return []byte{}
	}
	return buffer
}

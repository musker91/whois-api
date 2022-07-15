package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetRootPath() (rootPath string) {
	rootPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(fmt.Sprintf("get project root path faild: %s", err))
	}
	return
}

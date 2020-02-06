package file_manager

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"
)

func CurrentExeDir() string {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	return cwdPath
}

func CurrentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		return ""
	}
	return strings.Replace(dir, "\\", "/", -1)
}

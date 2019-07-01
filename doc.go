package file_manager

import (

	"runtime"
	"path"
)

func CurrentDir()string  {
	_, filename, _, ok := runtime.Caller(1)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	}  else  {
		cwdPath = "./"
	}
	return cwdPath
}
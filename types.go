package file_manager

import (
	"time"
	"os"
)

type FileInfo struct {
	Name  string
	Size  int64
	ModifyTime time.Time
	IsDir   bool
	FileMode os.FileMode

	FileInfo FileInfo
}

func makeFileInfo(info os.FileInfo)FileInfo  {
	return FileInfo{
		Name:info.Name(),
		Size:info.Size(),
		ModifyTime:info.ModTime(),
		IsDir:info.IsDir(),
		FileMode:info.Mode(),
	}
}
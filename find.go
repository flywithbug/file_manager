package file_manager




import (
	"path/filepath"
	"os"
)

func GetFileList(root string)([]string,error)  {
	list := make([]string,0)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		list = append(list,path)
		return nil
	})
	if err != nil {
		return nil,err
	}
	return list,nil
}
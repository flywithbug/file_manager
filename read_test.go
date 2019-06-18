package file_manager

import (
	"testing"
	"fmt"
)

func TestGetFileList(t *testing.T) {
	list,err := GetFileList("./")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(list)
}

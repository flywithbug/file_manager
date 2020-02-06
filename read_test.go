package file_manager

import (
	"fmt"
	"testing"
)

func TestGetFileList(t *testing.T) {
	list, err := GetFileList("./")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
}

func TestReadAll(t *testing.T) {
	b, err := ReadAll("./read.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}

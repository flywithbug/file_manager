package file_manager

import (
	"testing"
	"fmt"
)

func TestPathExists(t *testing.T) {
	e,err := PathExists("./")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(e)
}

func TestCreateFile(t *testing.T) {
	err := createPath("./test1/test2/test3/1.h")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestCreatePath(t *testing.T) {
	err := CreatePath("./test1/test2/test3/test4/1.h")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestRemoveAll(t *testing.T) {
	err := RemoveAll("./test1")
	if err != nil {
		fmt.Println(err.Error())
	}
}
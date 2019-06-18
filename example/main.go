package main

import (
	"github.com/flywithbug/file_manager"
	"fmt"
)

func main()  {
	//e,err := file_manager.PathExists("./test")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(e)

	err := file_manager.WriteFileString("./test1/1.m","test.a.com.ccc",true)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = file_manager.WriteFile("./test1/1.h",[]byte("test.a.com.cc"),true)
	if err != nil {
		fmt.Println(err.Error())
	}


	//file_manager.RemoveAll("./test1")
}

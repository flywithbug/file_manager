package main

import (
	"fmt"
	"sync"
	"github.com/flywithbug/file_manager"
)

func main()  {
	//e,err := file_manager.PathExists("./test")
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(e)

	//err := file_manager.WriteFileString("./test1/1.m","test.a.com.ccc",true)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//err = file_manager.WriteFile("./test1/1.h",[]byte("test.a.com.cc"),true)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}


	//file_manager.RemoveAll("./test1")

	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",

	}
	count := 0
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		go func(url string) {
			// Launch a goroutine to fetch the URL.
			defer wg.Done()
			count++
			err := file_manager.WriteFileString(fmt.Sprintf("./test1/%d.m", count),url,true)
			if err != nil {
				fmt.Println(err.Error())
			}
			err = file_manager.WriteFile(fmt.Sprintf("./test1/%d.m", count),[]byte(url),true)
			if err != nil {
				fmt.Println(err.Error())
			}
		}(url)
	}
	// Wait for all goroutines to finish.
	wg.Wait()
	fmt.Println("Game Over")



}

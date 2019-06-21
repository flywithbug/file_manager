package main

import (
	"sync"
	"github.com/flywithbug/file_manager"
	"fmt"
)

func main()  {
	//writeFileTest()
	readFileList()
}

func readFileList()  {
	list,err := file_manager.GetFileList("./")
	if err != nil{
		fmt.Println(err)
	}
	//fmt.Println(list)
	for _,v := range list{
		if !v.IsDir {
			fmt.Println(v.Name)
		}
	}
}

func writeFileTest()  {

	var wg sync.WaitGroup
	var urls = []string{
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
